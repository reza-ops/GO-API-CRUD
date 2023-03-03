package book_migration

import (
	"example/server/db"
	"log"
	"time"
)

type Books struct {
	ID          int       `json:"id" gorm:"primary key"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       string    `json:"price"`
	IsActive    string    `json:"is_active"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at"`
}

func MigrateBooks() {
	db.DataBase.AutoMigrate(&Books{})
	log.Println("Book Migration Success")
}
