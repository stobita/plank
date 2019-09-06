package controller

type postBoardsRequestBody struct {
	Title string `json:"title" binding:"required"`
}
