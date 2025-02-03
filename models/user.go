package models

import (
	"errors"
	"event-booking/db"
	"event-booking/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u *User) Save() error {
	query := "INSERT INTO users (email, password) VALUES (?, ?)"

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	hashedPass, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.Email, hashedPass)
	if err != nil {
		return err
	}

	userId, err := result.LastInsertId()
	u.ID = userId

	return err
}

func (u *User) ValidateCredentials() error {
	query := "SELECT password, id FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, u.Email)

	var retrievedPass string
	err := row.Scan(&retrievedPass, &u.ID)
	if err != nil {
		return err
	}

	if utils.CheckPasswordHash(u.Password, retrievedPass) {
		return nil
	}

	return errors.New("invalid credentials")
}
