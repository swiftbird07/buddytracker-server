package controller

import (
	"time"

	"github.com/swiftbird07/buddytracker-server/utils"
)

// Temp
var users []User

type User struct {
	Id         string
	udid       string
	token      string
	Name       string
	FriendCode FriendCode
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type FriendCode struct {
	Code      string
	CreatedAt time.Time
}

func NewUser(udid string, name string) (User, error) {
	for _, user := range users {
		if user.udid == udid {
			return user, nil
		}
	}

	newUser := User{
		Id:         utils.GenerateId(10, "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"),
		udid:       udid,
		token:      "1234", // utils.GenerateId(20, "0123456789abcdefghijklmnopqrstuvwxyz")
		Name:       name,
		FriendCode: FriendCode{Code: utils.GenerateId(8, "123456789ABCDEFGHIJKLMNPQRSTUVWXYZ"), CreatedAt: time.Time{}},
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	users = append(users, newUser)

	return newUser, nil
}

func (u *User) ChangeName(name string) {
	u.Name = name
	u.UpdatedAt = time.Now()
}

func (u *User) GetToken() string {
	return u.token
}

func ValidToken(token string) (bool, User) {
	for _, user := range users {
		if user.token == token {
			return true, user
		}
	}

	return false, User{}
}
