package entity

import "gorm.io/gorm"

// Album represents data about a record album.
type Album struct {
	gorm.Model
	Title    string
	Artist   string
	Price    float32
	Quantity int
}
