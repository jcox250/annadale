package usecases

import "tech-test-jc/domain"

type PersonInteractor struct {
	PersonRepository domain.PersonRepository
}

func (p PersonInteractor) People() ([]domain.Person, error) {
	return p.PersonRepository.GetPeople()
}

func (p PersonInteractor) AddPerson(person domain.Person) (int64, error) {
	return p.PersonRepository.AddPerson(person)
}
