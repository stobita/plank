package controller

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stobita/plank/internal/event"
	"github.com/stobita/plank/internal/model"
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

func (c *Controller) GetBoardLabels() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		boardID, err := strconv.Atoi(ctx.Param("boardID"))
		if err != nil {
			log.Print(err)
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid params boardId"})
			return
		}
		labels, err := c.inputPort.GetBoardLabels(uint(boardID))
		if err != nil {
			log.Println(err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}
		res, err := presenter.GetLabelsResponse(labels)
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
		var reqBody postBoardsSectionsCardsRequestBody
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
		limitTime := time.Unix(int64(reqBody.LimitTime), 0)
		input := usecase.CreateCardInput{
			Name:        reqBody.Name,
			Description: reqBody.Description,
			SectionID:   uint(sectionID),
			Labels:      reqBody.Labels,
			LimitTime:   &limitTime,
			Images:      reqBody.Images,
		}
		result, err := c.inputPort.CreateCard(input)
		if err != nil {
			log.Print(err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}
		res, err := presenter.GetCardResponse(result)
		if err != nil {
			log.Print(err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}
		ctx.JSON(http.StatusOK, res)
	}
}

func (c *Controller) GetLabelSections() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		labelID, err := strconv.Atoi(ctx.Param("labelID"))
		if err != nil {
			log.Print(err)
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid params"})
			return
		}
		boardID, err := strconv.Atoi(ctx.Param("boardID"))
		if err != nil {
			log.Print(err)
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid params"})
			return
		}
		sections, err := c.inputPort.GetLabelSections(uint(boardID), uint(labelID))
		if err != nil {
			log.Println(err)
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

func (c *Controller) GetBoardsSections() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		boardID, err := strconv.Atoi(ctx.Param("boardID"))
		if err != nil {
			log.Print(err)
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid params"})
			return
		}
		word := ctx.Query("word")
		var sections []*model.Section
		if word == "" {
			sections, err = c.inputPort.GetBoardSections(uint(boardID))

		} else {
			sections, err = c.inputPort.SearchBoardsSections(uint(boardID), word)

		}
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

func (c *Controller) PutBoardsSectionsCards() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var reqBody putBoardsSectionsCardsRequestBody
		if err := ctx.ShouldBindJSON(&reqBody); err != nil {
			log.Print(err)
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid params"})
			return
		}
		cardID, err := strconv.Atoi(ctx.Param("cardID"))
		if err != nil {
			log.Print(err)
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid params"})
			return
		}
		limitTime := time.Unix(int64(reqBody.LimitTime), 0)
		input := usecase.UpdateCardInput{
			Name:        reqBody.Name,
			Description: reqBody.Description,
			Labels:      reqBody.Labels,
			LimitTime:   &limitTime,
		}

		result, err := c.inputPort.UpdateCard(cardID, input)
		if err != nil {
			log.Print(err)
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

func (c *Controller) ReorderSection() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var reqBody reorderSectionRequestBody
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

		if err := c.inputPort.ReorderSectionPosition(uint(sectionID), reqBody.Position); err != nil {
			log.Print(err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}
		ctx.JSON(http.StatusOK, nil)
	}
}

func (c *Controller) ReorderCard() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var reqBody reorderCardRequestBody
		if err := ctx.ShouldBindJSON(&reqBody); err != nil {
			log.Print(err)
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid params"})
			return
		}
		cardID, err := strconv.Atoi(ctx.Param("cardID"))
		if err != nil {
			log.Print(err)
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid params"})
			return
		}

		if err := c.inputPort.ReorderCardPosition(uint(cardID), reqBody.Position); err != nil {
			log.Print(err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}
		ctx.JSON(http.StatusOK, nil)
	}
}

func (c *Controller) MoveCard() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var reqBody moveCardRequestBody
		if err := ctx.ShouldBindJSON(&reqBody); err != nil {
			log.Print(err)
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid params"})
			return
		}
		cardID, err := strconv.Atoi(ctx.Param("cardID"))
		if err != nil {
			log.Print(err)
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid params"})
			return
		}

		if err := c.inputPort.MoveCardPosition(uint(cardID), reqBody.Position, reqBody.DestinationSectionID); err != nil {
			log.Print(err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}
		ctx.JSON(http.StatusOK, nil)
	}
}

func (c *Controller) DeleteBoardsSectionsCards() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		cardID, err := strconv.Atoi(ctx.Param("cardID"))
		if err != nil {
			log.Print(err)
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid params"})
			return
		}
		if err := c.inputPort.DeleteCard(cardID); err != nil {
			log.Print(err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}
		ctx.JSON(http.StatusNoContent, nil)
	}
}

func (c *Controller) PutBoardsSections() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var reqBody putBoardsSectionsRequestBody
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
		input := usecase.UpdateSectionInput{
			Name: reqBody.Name,
		}
		result, err := c.inputPort.UpdateSection(sectionID, input)
		if err != nil {
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

func (c *Controller) DeleteBoadsSections() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		sectionID, err := strconv.Atoi(ctx.Param("sectionID"))
		if err != nil {
			log.Print(err)
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid params"})
			return
		}
		if err := c.inputPort.DeleteSection(sectionID); err != nil {
			log.Print(err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}
		ctx.JSON(http.StatusNoContent, nil)
	}
}

func (c *Controller) PutBoards() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var reqBody putBoardsRequestBoady
		if err := ctx.ShouldBindJSON(&reqBody); err != nil {
			log.Print(err)
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid params"})
			return
		}
		boardID, err := strconv.Atoi(ctx.Param("boardID"))
		if err != nil {
			log.Print(err)
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid params"})
			return
		}
		input := usecase.UpdateBoardInput{
			Name: reqBody.Name,
		}
		result, err := c.inputPort.UpdateBoard(boardID, input)
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

func (c *Controller) DeleteBoads() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		boardID, err := strconv.Atoi(ctx.Param("boardID"))
		if err != nil {
			log.Print(err)
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid params"})
			return
		}
		if err := c.inputPort.DeleteBoard(boardID); err != nil {
			log.Print(err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}
		ctx.JSON(http.StatusNoContent, nil)
	}
}

func (c *Controller) AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// sid, err := ctx.Cookie("sid")
		// if err != nil {
		// 	log.Print(err)
		// 	ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid params"})
		// 	ctx.Abort()
		// 	return
		// }
		// if err := c.inputPort.GetUserSession(sid); err != nil {
		// 	log.Print(err)
		// 	ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		// 	ctx.Abort()
		// 	return
		// }
		// ctx.Next()
	}
}
