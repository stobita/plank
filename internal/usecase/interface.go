package usecase

import "github.com/stobita/plank/internal/model"

// InputPort ...
type InputPort interface {
	GetAllBoards() ([]*model.Board, error)
	CreateBoard(CreateBoardInput) (*model.Board, error)
	CreateSection(CreateSectionInput) (*model.Section, error)
	CreateCard(CreateCardInput) (*model.Card, error)
	GetBoardSections(boardID uint) ([]*model.Section, error)

	AddEventClient(EventClient)
	RemoveEventClient(EventClient)

	UpdateBoard(id int, input UpdateBoardInput) (*model.Board, error)
	DeleteBoard(id int) error

	UpdateCard(id int, input UpdateCardInput) (*model.Card, error)
	MoveCardPosition(id uint, input MoveCardPositionInput) error
	DeleteCard(id int) error

	UpdateSection(id int, input UpdateSectionInput) (*model.Section, error)
	DeleteSection(id int) error
}

// Repository ...
type Repository interface {
	GetAllBoards() ([]*model.Board, error)
	GetBoard(id uint) (*model.Board, error)
	SaveNewBoard(*model.Board) error
	SaveBoard(*model.Board) error
	DeleteBoard(*model.Board) error

	GetSection(id uint) (*model.Section, error)
	SaveNewSection(*model.Section) error
	SaveSection(*model.Section) error
	DeleteSection(*model.Section) error

	GetBoardSectionsWithCards(*model.Board) ([]*model.Section, error)

	GetCard(id uint) (*model.Card, error)
	DeleteCard(*model.Card) error
	SaveCard(*model.Card) error
	MoveCardPosition(id uint, prevID uint, targetSectionID uint) error

	SaveNewCard(*model.Card) error
}

type eventBroker interface {
	AddClient(EventClient)
	RemoveClient(EventClient)

	PushAddCardEvent(*model.Board) error
}

type EventClient interface{}
