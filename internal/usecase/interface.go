package usecase

import "github.com/stobita/plank/internal/model"

type repository interface {
	GetAllBoards() ([]*model.Board, error)
	GetBoardByID(id uint) (*model.Board, error)
	SaveNewBoard(*model.Board) error

	GetSectionByID(id uint) (*model.Section, error)
	GetBoardSectionsWithCards(*model.Board) ([]*model.Section, error)
	SaveNewSection(*model.Section) error

	GetCard(id uint) (*model.Card, error)
	DeleteCard(*model.Card) error
	SaveCard(*model.Card) error

	SaveNewCard(*model.Card) error
}

type eventBroker interface {
	AddClient(EventClient)
	RemoveClient(EventClient)

	PushAddCardEvent(*model.Board) error
}

type EventClient interface{}

type InputPort interface {
	GetAllBoards() ([]*model.Board, error)
	CreateBoard(CreateBoardInput) (*model.Board, error)
	CreateSection(CreateSectionInput) (*model.Section, error)
	CreateCard(CreateCardInput) (*model.Card, error)
	GetBoardSections(boardID uint) ([]*model.Section, error)

	AddEventClient(EventClient)
	RemoveEventClient(EventClient)

	UpdateCard(id int, input UpdateCardInput) (*model.Card, error)
	DeleteCard(id int) error
}
