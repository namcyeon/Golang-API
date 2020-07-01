package models

import (
	"github.com/jinzhu/gorm"
)

type Food struct {
	ID     uint   `json:"id" gorm:"primary_key"`
	Name  string `json:"name"`
	Price double `json:"price"`
	DiscountPrice double `json:"discount_price"`
}