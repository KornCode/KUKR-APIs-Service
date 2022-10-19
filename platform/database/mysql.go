package database

import (
	"fmt"

	config "github.com/KornCode/KUKR-APIs-Service/pkg/configs"
	"github.com/KornCode/KUKR-APIs-Service/pkg/logs"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewConnectionMySQLDB(conf config.MySQLDB) (*gorm.DB, error) {
	sql_dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.Username,
		conf.Password,
		conf.Host,
		conf.Port,
		conf.DBName,
	)

	sql_db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: sql_dns,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
		DryRun: false,
	})
	if err != nil {
		logs.Error("MySQLDB is Not connected")

		return nil, err
	}

	fmt.Println("MySQLDB is Connected")

	sql_db.AutoMigrate(publishMySQL{})

	return sql_db, err
}
