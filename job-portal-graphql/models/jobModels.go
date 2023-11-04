package models

import "gorm.io/gorm"

type Job struct {
	gorm.Model
	Company   Company `json:"-" gorm:"foreignKey:company_id"`
	CompanyId uint64    `json:"Company_id"`
	Title     string  `json:"Title"`
	Desc      string  `json:"Desc"`
}