package usecases

type Repo interface {
	Get(string)
	GetAll()
}
