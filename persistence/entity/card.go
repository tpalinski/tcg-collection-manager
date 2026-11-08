package entity

import "gorm.io/gorm"

type Card struct {
	gorm.Model
  card_id string
  name string
  image_url string
}
