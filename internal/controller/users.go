package controller

import (
	"time"

	"github.com/swiftbird07/buddytracker-server/utils"
)

type User struct {
	Id         string
	Name       string
	FriendCode FriendCode
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type FriendCode struct {
	Code      string
	CreatedAt time.Time
}

func NewUser(name string) User {
	return User{
		Id:   utils.GenerateId(10, "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"),
		Name: name,
		FriendCode: FriendCode{
			Code:      utils.GenerateId(8, "123456789ABCDEFGHIJKLMNPQRSTUVWXYZ"),
			CreatedAt: time.Time{},
		},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (u *User) ChangeName(name string) {
	u.Name = name
	u.UpdatedAt = time.Now()
}
