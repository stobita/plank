package controller

type postBoardsRequestBody struct {
	Name string `json:"name" binding:"required"`
}

type postBoardsSectionsRequestBody struct {
	Name string `json:"name" binding:"required"`
}

type postBoardsSectionsCardsRequestBody struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}

type putBoardsSectionsCardsPositionRequestBody struct {
	Position int `json:"position"`
}

type putBoardsRequestBoady = postBoardsRequestBody

type putBoardsSectionsCardsRequestBody = postBoardsSectionsCardsRequestBody

type putBoardsSectionsRequestBody = postBoardsSectionsRequestBody
