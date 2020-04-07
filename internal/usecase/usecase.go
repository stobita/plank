package usecase

type usecase struct {
	repository  Repository
	eventBroker eventBroker
}

// New ...
func New(r Repository, e eventBroker) *usecase {
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
