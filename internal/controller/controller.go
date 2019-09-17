package controller

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/stobita/plank/internal/event"
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

func (c *Controller) SSESubscribe() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		flusher, ok := ctx.Writer.(http.Flusher)

		if !ok {
			http.Error(ctx.Writer, "Streaming unsupported!", http.StatusInternalServerError)
			return
		}

		ctx.Writer.Header().Set("Content-Type", "text/event-stream")
		ctx.Writer.Header().Set("Cache-Control", "no-cache")
		ctx.Writer.Header().Set("Connection", "keep-alive")

		sendChan := make(chan []byte)
		client := &event.Client{SendChannel: sendChan}

		c.inputPort.AddEventClient(client)
		defer c.inputPort.RemoveEventClient(client)

		for {
			select {
			case msg := <-sendChan:
				log.Println("receive")
				fmt.Fprintf(ctx.Writer, "data: %s\n\n", msg)
				flusher.Flush()
			case <-ctx.Request.Context().Done():
				log.Println("request context done")
				close(sendChan)
				return
			}
		}
	}
}
