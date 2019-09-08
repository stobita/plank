package usecase

import (
	"errors"

	"github.com/stobita/plank/internal/model"
)

type CreateBoardInput struct {
	Name string
}

type CreateSectionInput struct {
	Name    string
	BoardID uint
}

type CreateCardInput struct {
	Name        string
	Description string
	SectionID   uint
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

func (u *usecase) GetBoardSections(boardID uint) ([]*model.Section, error) {
	board, err := u.repository.GetBoardByID(boardID)
	if err != nil {
		return nil, err
	}
	if board == nil {
		return nil, errors.New("board not found")
	}
	sections, err := u.repository.GetBoardSectionsWithCards(board)
	if err != nil {
		return nil, err
	}
	return sections, nil
}

func (u *usecase) CreateSection(input CreateSectionInput) (*model.Section, error) {
	board, err := u.repository.GetBoardByID(input.BoardID)
	if err != nil {
		return nil, err
	}
	if board == nil {
		return nil, errors.New("Invalid board id")
	}
	m := &model.Section{
		Name:  input.Name,
		Board: board,
	}
	if err := u.repository.SaveNewSection(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (u *usecase) CreateCard(input CreateCardInput) (*model.Card, error) {
	section, err := u.repository.GetSectionByID(input.SectionID)
	if err != nil {
		return nil, err
	}
	if section == nil {
		return nil, errors.New("Invalid section id")
	}
	m := &model.Card{
		Name:        input.Name,
		Description: input.Description,
		Section:     section,
	}
	if err := u.repository.SaveNewCard(m); err != nil {
		return nil, err
	}
	return m, nil
}
