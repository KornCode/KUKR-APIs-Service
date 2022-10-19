package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	config "github.com/KornCode/KUKR-APIs-Service/pkg/configs"
	"github.com/KornCode/KUKR-APIs-Service/pkg/logs"
	route "github.com/KornCode/KUKR-APIs-Service/pkg/routes"
	"github.com/KornCode/KUKR-APIs-Service/platform/cache"
	"github.com/KornCode/KUKR-APIs-Service/platform/database"
	"github.com/KornCode/KUKR-APIs-Service/platform/web"
)

func init() {
	logs.InitZapLog()

	// liveness probe
	_, err := os.Create("/tmp/live")
	if err != nil {
		logs.Error(err)
	}
}

func main() {
	confs := config.NewConfigs()

	rd_cache, err := cache.NewConnectionRedisDB(confs.RedisDB)
	if err != nil {

	}

	sql_db, err := database.NewConnectionMySQLDB(confs.MySQLDB)
	if err != nil {
		panic(err)
	}

	app := web.NewFiberApp(confs.Server)
	app_api_v1 := app.Group("/v1/kukr")

	route.SetupPublishRoutes(app_api_v1, sql_db, rd_cache)

	go func() {
		if err := web.ListenFiberApp(confs.Server); err != nil {
			panic(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	_ = <-c
	_ = web.ShutdownFiberApp()
	_ = cache.CloseConnectionRedisDB()

	os.Remove("/tmp/live")

	fmt.Println("Graceful shutdown")
}
