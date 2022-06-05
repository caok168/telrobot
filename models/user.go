package models

import "telrobot/util/common"

type User struct {
	common.Model

	Uuid                 string `json:"uuid" gorm:"not null"`
	Name                 string `json:"name" gorm:"not null;unique_index"`
	Email                string `json:"email" gorm:"not null"`
	Phone                string `json:"phone" gorm:"not null"`
	Password             string `json:"password" gorm:"not null"`
	PasswordConfirmation string `json:"password_confirmation" gorm:"not null"`
}

type ResultCount struct {
	Count int `json:"count"`
}
