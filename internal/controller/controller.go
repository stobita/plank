package controller

import (
	"log"
	"net/http"
	"strconv"

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
			Name: reqBody.Name,
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

func (c *Controller) PostBoardsSections() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var reqBody postBoardsSectionsRequestBody
		if err := ctx.Bind(&reqBody); err != nil {
			log.Print(err)
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid params"})
			return
		}
		boardID, err := strconv.Atoi(ctx.Param("boardID"))
		if err != nil {
			log.Print(err)
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid params boardId"})
			return
		}
		input := usecase.CreateSectionInput{
			Name:    reqBody.Name,
			BoardID: uint(boardID),
		}
		result, err := c.inputPort.CreateSection(input)
		if err != nil {
			log.Print(err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}
		res, err := presenter.GetSectionResponse(result)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}
		ctx.JSON(http.StatusOK, res)
	}
}

func (c *Controller) PostBoardsSectionsCards() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var reqBody postBoardsSectionsCardRequestBody
		if err := ctx.ShouldBindJSON(&reqBody); err != nil {
			log.Print(err)
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid params"})
			return
		}
		sectionID, err := strconv.Atoi(ctx.Param("sectionID"))
		if err != nil {
			log.Print(err)
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid params"})
			return
		}
		input := usecase.CreateCardInput{
			Name:        reqBody.Name,
			Description: reqBody.Description,
			SectionID:   uint(sectionID),
		}
		result, err := c.inputPort.CreateCard(input)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}
		res, err := presenter.GetCardResponse(result)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}
		ctx.JSON(http.StatusOK, res)
	}
}

func (c *Controller) GetBoardsSections() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		boardID, err := strconv.Atoi(ctx.Param("boardID"))
		if err != nil {
			log.Print(err)
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid params"})
			return
		}
		sections, err := c.inputPort.GetBoardSections(uint(boardID))
		if err != nil {
			log.Print(err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}
		res, err := presenter.GetSectionsResponse(sections)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}
		ctx.JSON(http.StatusOK, res)
	}
}
