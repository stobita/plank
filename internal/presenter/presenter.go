package presenter

import "github.com/stobita/plank/internal/model"

type boardJSON struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
}

type listJSON struct {
	Items []interface{} `json:"items"`
}

func GetBoardsResponse(model []*model.Board) (listJSON, error) {
	json := listJSON{Items: []interface{}{}}
	for _, v := range model {
		json.Items = append(json.Items, &boardJSON{
			ID:    v.ID,
			Title: v.Title,
		})
	}
	return json, nil
}

func GetBoardResponse(model *model.Board) (boardJSON, error) {
	return boardJSON{
		ID:    model.ID,
		Title: model.Title,
	}, nil
}
