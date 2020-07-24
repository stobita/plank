package usecase

import "github.com/stobita/plank/internal/model"

// InputPort ...
type InputPort interface {
	GetAllBoards() ([]*model.Board, error)
	CreateBoard(CreateBoardInput) (*model.Board, error)
	CreateSection(CreateSectionInput) (*model.Section, error)
	CreateCard(CreateCardInput) (*model.Card, error)
	GetBoardSections(boardID uint) ([]*model.Section, error)
	GetLabelSections(boardID uint, labelID uint) ([]*model.Section, error)
	SearchBoardsSections(boardID uint, word string) ([]*model.Section, error)

	AddEventClient(EventClient)
	RemoveEventClient(EventClient)

	UpdateBoard(id int, input UpdateBoardInput) (*model.Board, error)
	DeleteBoard(id int) error

	UpdateCard(id int, input UpdateCardInput) (*model.Card, error)

	ReorderCardPosition(id uint, position uint) error
	MoveCardPosition(id uint, position uint, destinationSectionID uint) error

	DeleteCard(id int) error

	UpdateSection(id int, input UpdateSectionInput) (*model.Section, error)
	DeleteSection(id int) error
	ReorderSectionPosition(id uint, position uint) error

	GetBoardLabels(boardID uint) ([]*model.Label, error)
}

// Repository ...
type Repository interface {
	GetAllBoards() ([]*model.Board, error)
	GetBoard(id uint) (*model.Board, error)
	SaveNewBoard(*model.Board) error
	SaveBoard(*model.Board) error
	DeleteBoard(*model.Board) error

	SaveNewLabel(*model.Label) error
	GetLabelByName(name string) (*model.Label, error)

	GetSection(id uint) (*model.Section, error)
	SaveNewSection(*model.Section) error
	SaveSection(*model.Section) error
	DeleteSection(*model.Section) error
	ReorderSectionPosition(id uint, position uint) error

	GetBoardSectionsWithCards(*model.Board) ([]*model.Section, error)
	SearchBoardSectionsWithCards(*model.Board, string) ([]*model.Section, error)

	GetCard(id uint) (*model.Card, error)
	DeleteCard(*model.Card) error
	SaveCard(*model.Card) error
	ReorderCardPosition(id uint, position uint) error
	MoveCardPosition(id uint, position uint, destinationSectionID uint) error

	SaveNewCard(*model.Card) error

	GetBoardLabels(boardID uint) ([]*model.Label, error)
	GetLabel(id uint) (*model.Label, error)
}

type eventBroker interface {
	AddClient(EventClient)
	RemoveClient(EventClient)

	PushAddCardEvent(*model.Board) error
}

type EventClient interface{}
