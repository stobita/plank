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

type reorderCardRequestBody struct {
	Position uint `json:"position"`
}

type moveCardRequestBody struct {
	Position             uint `json:"position"`
	DestinationSectionID uint `json:"destinationSectionId"`
}

type putBoardsRequestBoady = postBoardsRequestBody

type putBoardsSectionsCardsRequestBody = postBoardsSectionsCardsRequestBody

type putBoardsSectionsRequestBody = postBoardsSectionsRequestBody
