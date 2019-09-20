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
	repository := repository.New(db)
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
		v1.GET("/sse", controller.SSESubscribe())
		v1.GET("/boards", controller.GetBoards())
		v1.POST("/boards", controller.PostBoards())
		v1.PUT("/boards/:boardID", controller.PutBoards())
		v1.DELETE("/boards/:boardID", controller.DeleteBoads())

		v1.GET("/boards/:boardID/sections", controller.GetBoardsSections())
		v1.POST("/boards/:boardID/sections", controller.PostBoardsSections())
		v1.PUT("/boards/:boardID/sections/:sectionID", controller.PutBoardsSections())
		v1.DELETE("/boards/:boardID/sections/:sectionID", controller.DeleteBoadsSections())

		v1.POST("/sections/:sectionID/cards", controller.PostBoardsSectionsCards())
		v1.PUT("/sections/:sectionID/cards/:cardID", controller.PutBoardsSectionsCards())
		v1.DELETE("/sections/:sectionID/cards/:cardID", controller.DeleteBoardsSectionsCards())
	}
	return r, nil
}
