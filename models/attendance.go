package models

import (
)

func init() {
	
}

type Attendance struct {
	Id       int 		`json:"id"`
	UserId   string		`json:"user_id"`
	attendance 	string	`json:"attendance"`
	TagId    int        `json:"department"`
	Gender   string		`json:"gender"`
	Age      string		`json:"age"`
	Address  string		`json:"address"`
	Email    string		`json:"email"`
	Role     int        `json:"role"`
	Salf     string     `json:"salf"`
}

