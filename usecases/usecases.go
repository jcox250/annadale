package usecases

import "tech-test-jc/domain"

// PersonInteractor does x
type PersonInteractor struct {
	PersonRepository domain.PersonRepository
}

// People will return an array of all the people stored in the database
func (p PersonInteractor) People() ([]domain.Person, error) {
	return p.PersonRepository.GetPeople()
}

// AddPerson foo
func (p PersonInteractor) AddPerson(person domain.Person) (int64, error) {
	return p.PersonRepository.AddPerson(person)
}
