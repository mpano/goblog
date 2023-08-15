package migration

import (
	"fmt"
	articleModels "goblog/internal/modules/article/models"
	userModels "goblog/internal/modules/user/models"
	"goblog/pkg/database"
	"log"
)

func Migrate() {
	db := database.Connection()
	err := db.AutoMigrate(&userModels.User{}, &articleModels.Article{})
	if err != nil {
		log.Fatal("error  can't migrate")
	}
	fmt.Printf("migration successfully")
}
