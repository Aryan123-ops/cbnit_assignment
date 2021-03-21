package models

import (
	"cbnit/config"
	"fmt"
)

type Signup struct {
	UserId   string `json:"user_id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Mobile   string `json:"mobile"`
}

func CreateNewUser(newuser Signup) (*Signup, error) {
	var err error
	if newuser.Name == "" || newuser.Username == "" || newuser.Password == "" || newuser.Mobile == "" {
		fmt.Println("Error")
	}
	db := config.GetDB()
	sqlStatement := `INSERT INTO users (name, username, password, mobile)
						VALUES ($1, $2, $3, $4)`
	_, err = db.Exec(sqlStatement, newuser.Name, newuser.Username, newuser.Password, newuser.Mobile)

	if err != nil {
		panic(err)
	}
	return &newuser, nil
}

func GetUers() ([]Signup, error) {
	var err error
	var users Signup
	var data []Signup
	db := config.GetDB()
	sqlStatement := `SELECT user_id, name, username, password, mobile from users`
	rows, err := db.Query(sqlStatement)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&users.UserId, &users.Name, &users.Username, &users.Password, &users.Mobile)
		if err != nil {
			fmt.Println(err)
		}
		res := Signup{
			UserId:   users.UserId,
			Name:     users.Name,
			Username: users.Username,
			Password: users.Password,
			Mobile:   users.Mobile}
		data = append(data, res)
	}
	return data, nil
}
