package usecase

type usecase struct {
	repository repository
}

func New(r repository) *usecase {
	return &usecase{
		repository: r,
	}
}
