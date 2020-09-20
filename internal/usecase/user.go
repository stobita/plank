package usecase

func (u *usecase) GetUserSession(sid string) error {
	return u.repository.GetUserSession(sid)
}
