package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title   string
	Company string
	Body    string
}
