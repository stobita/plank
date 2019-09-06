package usecase

import "github.com/stobita/plank/internal/model"

type repository interface {
	GetAllBoards() ([]*model.Board, error)
	SaveNewLink(*model.Board) error
}

type InputPort interface {
	GetAllBoards() ([]*model.Board, error)
	CreateBoard(CreateBoardInput) (*model.Board, error)
}
