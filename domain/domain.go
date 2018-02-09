package domain

type PersonRepository interface {
	GetPeople() ([]Person, error)
	AddPerson(Person) (int64, error)
}

type Person struct {
	ID      int    `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	Age     int    `json:"age,omitempty"`
	Email   string `json:"email,omitempty"`
	Balance int    `json:"balance,omitempty"`
	Address string `json:"address,omitempty"`
}
