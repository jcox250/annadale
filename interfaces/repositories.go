package interfaces

import (
	"database/sql"
	"log"
	"tech-test-jc/domain"
)

type DBHandler interface {
	Query(string) (*sql.Rows, error)
	Execute(string, ...interface{}) (int64, error)
}

type Row interface {
	Scan(interface{})
	Next() bool
}

type DBRepo struct {
	dbHandlers map[string]DBHandler
	dbHandler  DBHandler
}

type PersonRepo DBRepo

// NewPersonRepo constructor
func NewPersonRepo(dbHandlers map[string]DBHandler) *PersonRepo {
	pr := new(PersonRepo)
	if pr.dbHandler = dbHandlers["DBPersonRepo"]; pr.dbHandler == nil {
		log.Fatal("dbHandler cannot be nil")
	}
	pr.dbHandlers = dbHandlers
	return pr
}

func (p *PersonRepo) GetPeople() ([]domain.Person, error) {
	var person domain.Person
	var people []domain.Person

	row, err := p.dbHandler.Query("call spPersonGetAll()")
	if err != nil {
		return nil, err
	}

	for row.Next() {
		err := row.Scan(
			&person.ID,
			&person.Name,
			&person.Age,
			&person.Email,
			&person.Balance,
			&person.Address,
		)
		if err != nil {
			return nil, err
		}
		people = append(people, person)
	}
	return people, nil
}

func (p *PersonRepo) AddPerson(person domain.Person) (int64, error) {
	return p.dbHandler.Execute("call spPersonAddPerson(?, ?, ?, ?, ?, ?)",
		person.ID,
		person.Name,
		person.Age,
		person.Email,
		person.Balance,
		person.Address,
	)
}
