package models

import (
	"errors"
	"fmt"

	"apilibrary.com/db"
	"apilibrary.com/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u *User) UserSave() error {
	query := `INSERT INTO users(email, password) VALUES(?, ?)`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.Email, hashedPassword)

	if err != nil {
		return err
	}
	userId, err := result.LastInsertId()

	u.ID = userId
	return err

}

func (u *User) ValidateCredentials() error {
	query := `SELECT user_id, password FROM users WHERE email = ?`

	row := db.DB.QueryRow(query, u.Email)
	fmt.Println("row is: ", row)

	var reterivedPassword string
	err := row.Scan(&u.ID, &reterivedPassword)

	if err != nil {
		return errors.New("Credentials Invalid")
	}

	passwordIsValid := utils.ValidatePasswordHash(u.Password, reterivedPassword)
	fmt.Print("this worked: ", u.Password, reterivedPassword)

	if !passwordIsValid {
		return errors.New("Credentials Invalid")
	}

	return nil

}
