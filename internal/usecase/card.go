package usecase

import (
	"database/sql"
	"errors"
	"time"

	"github.com/stobita/plank/internal/model"
)

type CreateCardInput struct {
	Name        string
	Description string
	SectionID   uint
	Labels      []string
	LimitTime   *time.Time
	Images      []string
}

type UpdateCardInput struct {
	Name        string
	Description string
	Labels      []string
	LimitTime   *time.Time
}

type ReorderCardPositionInput struct {
	TargetSectionID uint
}

func (u *usecase) CreateCard(input CreateCardInput) (*model.Card, error) {
	section, err := u.repository.GetSection(input.SectionID)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	if section == nil {
		return nil, errors.New("Invalid section id")
	}

	labels := make([]*model.Label, len(input.Labels))
	for i, v := range input.Labels {
		label, err := u.repository.GetLabelByName(v)
		if err != nil && err != sql.ErrNoRows {
			return nil, err
		}
		if label == nil {
			label = &model.Label{
				Name:    v,
				BoardID: section.Board.ID,
			}
			if err := u.repository.SaveNewLabel(label); err != nil {
				return nil, err
			}
		}
		labels[i] = label
	}

	m := &model.Card{
		Name:        input.Name,
		Description: input.Description,
		Section:     section,
		Labels:      labels,
		LimitTime:   input.LimitTime,
	}

	if err := u.repository.SaveNewCard(m); err != nil {
		return nil, err
	}

	for _, v := range input.Images {
		result, err := u.repository.UploadCardFile(v)
		if err != nil {
			return nil, err
		}
		if err := u.repository.SaveCardImage(m.ID, result); err != nil {
			return nil, err
		}
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
	card.LimitTime = input.LimitTime

	labels := make([]*model.Label, len(input.Labels))
	for i, v := range input.Labels {
		label, err := u.repository.GetLabelByName(v)
		if err != nil && err != sql.ErrNoRows {
			return nil, err
		}
		if label == nil {
			label = &model.Label{
				Name:    v,
				BoardID: card.Section.Board.ID,
			}
			if err := u.repository.SaveNewLabel(label); err != nil {
				return nil, err
			}
		}
		labels[i] = label
	}
	card.Labels = labels

	return card, u.repository.SaveCard(card)
}

func (u *usecase) ReorderCardPosition(id uint, position uint) error {
	if err := u.repository.ReorderCardPosition(id, position); err != nil {
		return err
	}
	return nil
}

func (u *usecase) MoveCardPosition(id uint, position uint, destinationSectionID uint) error {
	if err := u.repository.MoveCardPosition(id, position, destinationSectionID); err != nil {
		return err
	}
	return nil
}

func (u *usecase) DeleteCard(id int) error {
	card, err := u.repository.GetCard(uint(id))
	if err != nil {
		return err
	}
	return u.repository.DeleteCard(card)
}
