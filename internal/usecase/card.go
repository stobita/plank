package usecase

import (
	"errors"

	"github.com/stobita/plank/internal/model"
)

type CreateCardInput struct {
	Name        string
	Description string
	SectionID   uint
}

type UpdateCardInput struct {
	Name        string
	Description string
}

type UpdateCardPositionInput struct {
	Position uint
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

func (u *usecase) UpdateCardPosition(id int, input UpdateCardPositionInput) (*model.Card, error) {
	card, err := u.repository.GetCard(uint(id))
	if err != nil {
		return nil, err
	}
	card.Position = input.Position

	if err := u.repository.SaveCardPosition(card); err != nil {
		return nil, err
	}
	return card, nil
}

func (u *usecase) DeleteCard(id int) error {
	card, err := u.repository.GetCard(uint(id))
	if err != nil {
		return err
	}
	return u.repository.DeleteCard(card)
}
