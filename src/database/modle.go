package database

import "time"

type User struct {
	Id int
	Uuid string
	Name string
	Email string
	Password string
	CreateTime time.Time
}

type Session struct {
	Id int
	Uuid string
	Name string
	Email string
	CreateTime time.Time
}
