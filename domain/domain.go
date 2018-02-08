package domain

// PersonRepository will do y
type PersonRepository interface {
	GetPeople() ([]Person, error)
}

// Person will do y
type Person struct {
	ID      int    `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	Age     int    `json:"age,omitempty"`
	Email   string `json:"email,omitempty"`
	Balance int    `json:"balance,omitempty"`
	Address string `json:"address,omitempty"`
}
