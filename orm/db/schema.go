package db

import (
	"time"
)

type Author struct {
	Name  string
	Email string
}

type Books struct {
	ID          uint   `gorm:"primaryKey"`
	BookName    string `gorm:"index"`
	Author      Author `gorm:"embedded;embeddedPrefix:author_"`
	PublishDate time.Time
}
