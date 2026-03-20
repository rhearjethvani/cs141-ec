// maps to java userthread

package main

import (
	"strconv"
	"path/filepath"
)

type User struct {
	ID int
}

func NewUser(id int) *User {
	return &User {
		ID: id,
	}
}

// filename the user reads from (ex. USER0)
func (u *User) InputFile() string {
	return "USER" + strconv.Itoa(u.ID)
}

// filepath resolution
func (u *User) InputPath() string {
	return filepath.Join("users", u.InputFile())
}