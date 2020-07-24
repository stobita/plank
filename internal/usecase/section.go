package usecase

import (
	"errors"

	"github.com/stobita/plank/internal/model"
)

type CreateSectionInput struct {
	Name    string
	BoardID uint
}

type UpdateSectionInput struct {
	Name string
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

func (u *usecase) SearchBoardsSections(boardID uint, word string) ([]*model.Section, error) {
	board, err := u.repository.GetBoard(boardID)
	if err != nil {
		return nil, err
	}
	if board == nil {
		return nil, errors.New("board not found")
	}
	sections, err := u.repository.SearchBoardSectionsWithCards(board, word)
	if err != nil {
		return nil, err
	}
	return sections, nil
}

func (u *usecase) GetLabelSections(boardID uint, labelID uint) ([]*model.Section, error) {
	label, err := u.repository.GetLabel(labelID)
	if err != nil {
		return nil, err
	}

	if label == nil {
		return nil, errors.New("label not found")
	}
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

	result := make([]*model.Section, len(sections))
	for i, section := range sections {
		cards := []*model.Card{}
		for _, card := range section.Cards {
			hasLabel := false
			for _, _label := range card.Labels {
				if _label.ID == label.ID {
					hasLabel = true
				}
			}
			if hasLabel {
				c := new(model.Card)
				c = card
				cards = append(cards, c)
			}
		}
		s := new(model.Section)
		s = section
		s.Cards = cards
		result[i] = s
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

func (u *usecase) UpdateSection(id int, input UpdateSectionInput) (*model.Section, error) {
	section, err := u.repository.GetSection(uint(id))
	if err != nil {
		return nil, err
	}
	section.Name = input.Name

	if err := u.repository.SaveSection(section); err != nil {
		return nil, err
	}
	return section, nil
}

func (u *usecase) DeleteSection(id int) error {
	section, err := u.repository.GetSection(uint(id))
	if err != nil {
		return err
	}
	return u.repository.DeleteSection(section)
}

func (u *usecase) ReorderSectionPosition(id uint, position uint) error {
	return u.repository.ReorderSectionPosition(id, position)
}
