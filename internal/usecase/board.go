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

type UpdateSectionInput struct {
	Name string
}

type CreateCardInput struct {
	Name        string
	Description string
	SectionID   uint
}

type UpdateCardInput struct {
	Name        string
	Description string
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
	board, err := u.repository.GetBoard(boardID)
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
	board, err := u.repository.GetBoard(input.BoardID)
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
	section, err := u.repository.GetSection(input.SectionID)
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
	u.eventBroker.PushAddCardEvent(section.Board)
	return m, nil
}

func (u *usecase) UpdateCard(id int, input UpdateCardInput) (*model.Card, error) {
	card, err := u.repository.GetCard(uint(id))
	if err != nil {
		return nil, err
	}
	card.Name = input.Name
	card.Description = input.Description

	return card, u.repository.SaveCard(card)
}

func (u *usecase) DeleteCard(id int) error {
	card, err := u.repository.GetCard(uint(id))
	if err != nil {
		return err
	}
	return u.repository.DeleteCard(card)
}

func (u *usecase) UpdateSection(id int, input UpdateSectionInput) (*model.Section, error) {
	section, err := u.repository.GetSection(uint(id))
	if err != nil {
		return nil, err
	}
	section.Name = input.Name

	return section, u.repository.SaveSection(section)
}

func (u *usecase) DeleteSection(id int) error {
	section, err := u.repository.GetSection(uint(id))
	if err != nil {
		return err
	}
	return u.repository.DeleteSection(section)
}
