package usecase

import "github.com/stobita/plank/internal/model"

type repository interface {
	GetAllBoards() ([]*model.Board, error)
	GetBoardByID(id uint) (*model.Board, error)
	SaveNewBoard(*model.Board) error

	GetSectionByID(id uint) (*model.Section, error)
	GetBoardSectionsWithCards(*model.Board) ([]*model.Section, error)
	SaveNewSection(*model.Section) error

	SaveNewCard(*model.Card) error
}

type InputPort interface {
	GetAllBoards() ([]*model.Board, error)
	CreateBoard(CreateBoardInput) (*model.Board, error)
	CreateSection(CreateSectionInput) (*model.Section, error)
	CreateCard(CreateCardInput) (*model.Card, error)
	GetBoardSections(boardID uint) ([]*model.Section, error)
}
