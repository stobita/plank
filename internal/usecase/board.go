package usecase

import "github.com/stobita/plank/internal/model"

type CreateBoardInput struct {
	Title string
}

func (u *usecase) GetAllBoards() ([]*model.Board, error) {
	return u.repository.GetAllBoards()
}

func (u *usecase) CreateBoard(input CreateBoardInput) (*model.Board, error) {
	m := &model.Board{
		Title: input.Title,
	}
	if err := u.repository.SaveNewLink(m); err != nil {
		return nil, err
	}
	return m, nil
}
