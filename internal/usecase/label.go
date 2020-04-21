package usecase

import "github.com/stobita/plank/internal/model"

func (u *usecase) GetBoardLabels(boardID uint) ([]*model.Label, error) {
	return u.repository.GetBoardLabels(boardID)
}
