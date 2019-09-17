package presenter

import "github.com/stobita/plank/internal/model"

type boardJSON struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type sectionJSON struct {
	ID    uint       `json:"id"`
	Name  string     `json:"name"`
	Cards []cardJSON `json:"cards"`
}

type cardJSON struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type addCardEvent struct {
	Type    string `json:"type"`
	BoardID uint   `json:"boardId"`
}

const addCardEventType = "addCardEvent"

func GetAddCardEvent(model *model.Board) (addCardEvent, error) {
	return addCardEvent{
		Type:    addCardEventType,
		BoardID: model.ID,
	}, nil
}

type listJSON struct {
	Items []interface{} `json:"items"`
}

func GetBoardsResponse(model []*model.Board) (listJSON, error) {
	json := listJSON{Items: []interface{}{}}
	for _, v := range model {
		json.Items = append(json.Items, &boardJSON{
			ID:   v.ID,
			Name: v.Name,
		})
	}
	return json, nil
}

func GetBoardResponse(model *model.Board) (boardJSON, error) {
	return boardJSON{
		ID:   model.ID,
		Name: model.Name,
	}, nil
}

func GetSectionResponse(model *model.Section) (sectionJSON, error) {
	return sectionJSON{
		ID:   model.ID,
		Name: model.Name,
	}, nil
}

func GetSectionsResponse(model []*model.Section) (listJSON, error) {
	json := listJSON{Items: []interface{}{}}
	for _, v := range model {
		var cards []cardJSON
		for _, card := range v.Cards {
			cards = append(cards, cardJSON{
				ID:          card.ID,
				Name:        card.Name,
				Description: card.Description,
			})
		}
		json.Items = append(json.Items, &sectionJSON{
			ID:    v.ID,
			Name:  v.Name,
			Cards: cards,
		})
	}
	return json, nil
}

func GetCardResponse(model *model.Card) (cardJSON, error) {
	return cardJSON{
		ID:          model.ID,
		Name:        model.Name,
		Description: model.Description,
	}, nil
}
