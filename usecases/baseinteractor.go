package usecases

type BaseRepo interface {
	GetResults() ([]string, error)
	GetFixtures() ([]string, error)
}

type BaseInteractor struct {
	repo BaseRepo
}

func NewBaseInteractor(repo BaseRepo) BaseInteractor {
	return BaseInteractor{
		repo: repo,
	}
}

func (b BaseInteractor) GetFixtures() ([]string, error) {
	return b.repo.GetFixtures()
}

func (b BaseInteractor) GetLatestResults() ([]string, error) {
	return b.repo.GetResults()
}
