package models

import (
	"errors"
	"fmt"
)

type User struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Mobile       string `json:"mobile"`
	Status       string `json:"status"`
	LastModified uint64 `json:"last_modified"` //utc
}

func (u *User) Validate() error {
	if u.Name == "" {
		return errors.New("invalid name")
	}

	if u.Email == "" {
		return errors.New("invalid email")
	}
	return nil
}

func (u *User) ToString() string {
	return fmt.Sprintf("ID:%d Name: %s Email: %s Mobile:%s Status: %s Last-Modified: %d\n", u.Id, u.Name, u.Email, u.Mobile, u.Status, u.LastModified)
}
