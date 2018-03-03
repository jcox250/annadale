package domain

type User struct {
	ID        string `json:"id,omitempty"`
	Firstname string `json:"firstname,omitempty"`
	Surname   string `json:"surname,omitempty"`
	Email     string `json:"email,omitempty"`
	Username  string `json:"username,omitempty"`
	Password  string `json:"password,omitempty"`
	Archive   bool   `json:"archive,omitempty"`
}
