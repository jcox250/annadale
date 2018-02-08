package domain

// PersonRepository will do y
type PersonRepository interface {
	GetPeople() ([]Person, error)
}

// Person will do y
type Person struct {
	Name string `json:"name,omitempty"`
}
