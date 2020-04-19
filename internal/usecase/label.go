package usecase

import "github.com/stobita/plank/internal/model"

func (u *usecase) GetLabels() ([]*model.Label, error) {
	return u.repository.GetLabels()
}
