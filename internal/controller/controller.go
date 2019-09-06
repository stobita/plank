package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/stobita/plank/internal/presenter"
	"github.com/stobita/plank/internal/usecase"
)

type Controller struct {
	inputPort usecase.InputPort
}

func New(i usecase.InputPort) *Controller {
	return &Controller{
		inputPort: i,
	}
}

func (c *Controller) GetBoards() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		boards, err := c.inputPort.GetAllBoards()
		if err != nil {
			log.Println(err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}
		res, err := presenter.GetBoardsResponse(boards)
		if err != nil {
			log.Println(err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}
		ctx.JSON(http.StatusOK, res)
	}
}

func (c *Controller) PostBoards() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var reqBody postBoardsRequestBody
		if err := ctx.ShouldBindJSON(&reqBody); err != nil {
			log.Print(err)
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid params"})
			return
		}
		input := usecase.CreateBoardInput{
			Title: reqBody.Title,
		}
		result, err := c.inputPort.CreateBoard(input)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}
		res, err := presenter.GetBoardResponse(result)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}
		ctx.JSON(http.StatusOK, res)
	}
}
