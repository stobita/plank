package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/stobita/plank/internal/controller"
	"github.com/stobita/plank/internal/event"
	"github.com/stobita/plank/internal/infrastructure"
	"github.com/stobita/plank/internal/repository"
	"github.com/stobita/plank/internal/usecase"
)

func Run() error {
	db, err := infrastructure.NewDBConn()
	if err != nil {
		return err
	}
	esClient, err := infrastructure.NewESClient()
	if err != nil {
		return err
	}
	redisClient := infrastructure.NewRedisClient()
	fileClient := infrastructure.NewS3Client()
	repository := repository.New(db, esClient, fileClient, redisClient)
	eventBroker := event.NewBroker()
	usecase := usecase.New(repository, eventBroker)
	controller := controller.New(usecase)

	eventBroker.Run()

	engine, err := getEngine(controller)
	if err != nil {
		return err
	}
	return engine.Run()
}

func getEngine(controller *controller.Controller) (*gin.Engine, error) {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders: []string{"Origin", "Content-Length", "Content-Type"},
		AllowOrigins: []string{"http://localhost:3000"},
	}))
	v1 := r.Group("/api/v1")
	{
		authorized := v1.Group("/")
		authorized.Use(controller.AuthMiddleware())
		{
			authorized.GET("/sse", controller.SSESubscribe())
			authorized.GET("/boards", controller.GetBoards())
			authorized.POST("/boards", controller.PostBoards())
			authorized.PUT("/boards/:boardID", controller.PutBoards())
			authorized.DELETE("/boards/:boardID", controller.DeleteBoads())

			authorized.GET("/boards/:boardID/sections", controller.GetBoardsSections())
			authorized.POST("/boards/:boardID/sections", controller.PostBoardsSections())
			authorized.PUT("/boards/:boardID/sections/:sectionID", controller.PutBoardsSections())
			authorized.DELETE("/boards/:boardID/sections/:sectionID", controller.DeleteBoadsSections())
			authorized.PUT("/boards/:boardID/sections/:sectionID/reorder", controller.ReorderSection())
			authorized.GET("/boards/:boardID/labels/:labelID/sections", controller.GetLabelSections())

			authorized.GET("/boards/:boardID/labels", controller.GetBoardLabels())

			authorized.POST("/sections/:sectionID/cards", controller.PostBoardsSectionsCards())
			authorized.PUT("/sections/:sectionID/cards/:cardID", controller.PutBoardsSectionsCards())
			authorized.PUT("/sections/:sectionID/cards/:cardID/reorder", controller.ReorderCard())
			authorized.PUT("/sections/:sectionID/cards/:cardID/move", controller.MoveCard())
			authorized.DELETE("/sections/:sectionID/cards/:cardID", controller.DeleteBoardsSectionsCards())
		}
	}
	return r, nil
}
