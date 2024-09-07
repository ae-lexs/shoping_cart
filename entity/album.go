package entity

import "gorm.io/gorm"

type Album struct {
	gorm.Model
	Title    string
	Artist   string
	Price    float32
	Quantity int
}
