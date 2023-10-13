package domain

import "time"

type User struct {
	ID         int
	Name       string
	Email      string
	Password   string
	CreateTime time.Time
}

func (User) TableName() string {
	return "user"
}
