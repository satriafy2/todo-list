package db

import (
	"fmt"
	"todo-list/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Conn *gorm.DB

func Init() {
	dsn := "host=localhost user=postgres password=root dbname=todo_list port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("DB Connection err: ", err)
	}

	Conn = connection
}

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&model.Todo{},
	)
	fmt.Println("Migration running?")
}
