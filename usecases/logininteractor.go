package usecases

import "github.com/jcox250/annadale/domain"

type LoginRepo interface {
}

type LoginInteractor struct {
	Repo LoginRepo
}

func NewLoginInteractor(repo LoginRepo) *LoginInteractor {
	return &LoginInteractor{
		Repo: repo,
	}
}

func (l LoginInteractor) VerifyLogin(user domain.User) (bool, error) {
	foo := domain.User{
		Username: "james",
		Password: "james",
	}

	if user.Username == foo.Username && user.Password == foo.Password {
		return true, nil
	}

	return false, nil
}
