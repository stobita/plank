package usecase

type usecase struct {
	repository  repository
	eventBroker eventBroker
}

func New(r repository, e eventBroker) *usecase {
	return &usecase{
		repository:  r,
		eventBroker: e,
	}
}

func (u *usecase) AddEventClient(e EventClient) {
	u.eventBroker.AddClient(e)
}

func (u *usecase) RemoveEventClient(e EventClient) {
	u.eventBroker.RemoveClient(e)
}
