package server

import (
	"github.com/gin-gonic/gin"
	"github.com/stobita/plank/internal/controller"
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
	usecase := usecase.New(repository)
	controller := controller.New(usecase)

	engine, err := getEngine(controller)
	if err != nil {
		return err
	}
	return engine.Run()
}

func getEngine(controller *controller.Controller) (*gin.Engine, error) {
	r := gin.Default()
	v1 := r.Group("/api/v1")
	{
		v1.GET("/boards", controller.GetBoards())
		v1.POST("/boards", controller.PostBoards())
	}
	return r, nil
}
