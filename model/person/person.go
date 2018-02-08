package person

import (
	"flexbdna/model"
	"fmt"
	"log"
)

type Person struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Email   string `json:"email"`
	Balance int    `json:"balance"`
	Address string `json:"address"`
}

func GetAll() ([]Person, error) {
	conn, err := model.ConnectToDB()
	if err != nil {
		log.Println("error connecting to db: ", err)
	}
	defer model.SafeClose(conn, &err)

	rows, err := conn.Query("call spPersonGetAll()")
	if err != nil {
		return nil, fmt.Errorf("error calling spPersonGetAll(): %v", err)
	}
	defer model.SafeClose(rows, &err)

	var person Person
	var people []Person
	for rows.Next() {
		err = rows.Scan(
			&person.Id,
			&person.Name,
			&person.Age,
			&person.Email,
			&person.Balance,
			&person.Address,
		)
		if err != nil {
			return nil, fmt.Errorf("error scanning rows into []Person: %v", err)
		}
		people = append(people, person)
	}
	return people, nil
}

func Create(p Person) (int64, error) {
	conn, err := model.ConnectToDB()
	if err != nil {
		log.Println("error connecting to db: ", err)
	}
	defer model.SafeClose(conn, &err)

	res, err := conn.Exec("call spPersonAdd(?,?,?,?,?,?)",
		p.Id,
		p.Name,
		p.Age,
		p.Email,
		p.Balance,
		p.Address,
	)
	if err != nil {
		return 0, fmt.Errorf("error calling spPersonAdd: %v", err)
	}
	return res.RowsAffected()
}
