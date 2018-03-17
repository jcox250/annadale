package usecases

type BaseRepo interface {
	GetLatestResults() ([]string, error)
}

type BaseInteractor struct {
	baseRepo BaseRepo
}

func NewBaseInteractor(repo BaseRepo) BaseInteractor {
	return BaseInteractor{
		baseRepo: repo,
	}
}

func (b BaseInteractor) GetFixtures() ([]string, error) {
	return []string{
		"1s - Banbridge",
		"2s - Antrim",
		"3s - Raphoe",
		"4s - Bangor",
	}, nil
}

func (b BaseInteractor) GetLatestResults() ([]string, error) {
	return []string{
		"1s - Instonians",
		"2s - Three Rock",
		"3s - North Down",
		"4s - Civil Service",
	}, nil
}
