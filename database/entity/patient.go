package entity

import "gorm.io/gorm"

type Patient struct {
	gorm.Model
	Name   string
	Age    int
	Gender string
}
