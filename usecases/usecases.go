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
