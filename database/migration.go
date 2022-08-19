package database

import (
	"fmt"
	"waystalk/models"
	"waystalk/pkg/mysql"
)

func RunMigration() {
	err := mysql.DB.AutoMigrate(&models.Book{})

	if err != nil {
		fmt.Println(err)
		panic("Migration Failed")
	}

	fmt.Println("Migration Success")
}
