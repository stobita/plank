package controller

type postBoardsRequestBody struct {
	Name string `json:"name" binding:"required"`
}

type postBoardsSectionsRequestBody struct {
	Name string `json:"name" binding:"required"`
}

type postBoardsSectionsCardRequestBody struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}

type putBoardsRequestBoady = postBoardsRequestBody

type putBoardsSectionsCardRequestBody = postBoardsSectionsCardRequestBody

type putBoardsSectionsRequestBody = postBoardsSectionsRequestBody
