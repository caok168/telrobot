package models

import "telrobot/util/common"

type User struct {
	common.Model

	Uuid                 string `json:"uuid" gorm:"not null"`
	Name                 string `json:"name" gorm:"not null"`
	Email                string `json:"email" gorm:"not null"`
	Phone                string `json:"phone" gorm:"not null"`
	Password             string `gorm:"not null"`
	PasswordConfirmation string `gorm:"not null"`
}

type ResultCount struct {
	Count int `json:"count"`
}
