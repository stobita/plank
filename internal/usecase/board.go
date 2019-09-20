package usecase

import (
	"github.com/stobita/plank/internal/model"
)

type CreateBoardInput struct {
	Name string
}

type UpdateBoardInput struct {
	Name string
}

func (u *usecase) GetAllBoards() ([]*model.Board, error) {
	return u.repository.GetAllBoards()
}

func (u *usecase) CreateBoard(input CreateBoardInput) (*model.Board, error) {
	m := &model.Board{
		Name: input.Name,
	}
	if err := u.repository.SaveNewBoard(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (u *usecase) UpdateBoard(id int, input UpdateBoardInput) (*model.Board, error) {
	board, err := u.repository.GetBoard(uint(id))
	if err != nil {
		return nil, err
	}
	board.Name = input.Name

	if err := u.repository.SaveBoard(board); err != nil {
		return nil, err
	}
	return board, nil
}

func (u *usecase) DeleteBoard(id int) error {
	board, err := u.repository.GetBoard(uint(id))
	if err != nil {
		return err
	}
	return u.repository.DeleteBoard(board)
}
