package models

import "gorm.io/gorm"

type IceCream struct {
	gorm.Model

	Flavor string `gorm:"not null;unique_index" json:"flavor"`
	Stock  int    `gorm:"not null" json:"stock"`
}
