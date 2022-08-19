package models

import "time"

type Book struct {
	ID          int       `json:"id"`
	Title       string    `json:"title" gorm:"type: varchar(255)"`
	Description string    `json:"desc" gorm:"type: varchar(255)"`
	Price       int       `json:"price" gorm:"type: varchar(255)"`
	CreatedAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"-"`
}
