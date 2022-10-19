package routes

import (
	handler "github.com/KornCode/KUKR-APIs-Service/app/handlers"
	repository "github.com/KornCode/KUKR-APIs-Service/app/repositories"
	service "github.com/KornCode/KUKR-APIs-Service/app/services"
	"github.com/go-redis/redis/v9"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/helmet/v2"
	"gorm.io/gorm"
)

func SetupPublishRoutes(r fiber.Router, sql_db *gorm.DB, rd_cache *redis.Client) {
	publishRepository := repository.NewPublishRepositoryDB(sql_db, rd_cache)
	publishService := service.NewPublishService(publishRepository)
	publishHandler := handler.NewPublishHandler(publishService)

	r.Use(
		cors.New(),
		helmet.New(),
	)
	publishRoute := r.Group("/publishes")

	publishRoute.Post("/crud/create",
		publishHandler.CreateOne,
	)
	publishRoute.Put("/crud/update",
		publishHandler.UpdateOneByPK,
	)
	publishRoute.Delete("/crud/delete",
		publishHandler.DeleteOneByPK,
	)

	publishRoute.Post("/crud/sync-datasource",
		publishHandler.SyncDataSource,
	)

	publishRoute.Get("/info/category-pub_year",
		publishHandler.GetByCategoryAndPubYear,
	)
	publishRoute.Get("/info/bibid",
		publishHandler.GetByBibid,
	)
	publishRoute.Get("/info/paginate",
		publishHandler.GetPaginateByOptions,
	)

}
