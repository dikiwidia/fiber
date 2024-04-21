package migration

import (
	"fmt"

	"github.com/dikiwidia/fiber/database"
	"github.com/dikiwidia/fiber/models/entity"
)

func RunMigration() {
	err := database.DB.AutoMigrate(&entity.User{})
	if err != nil {
		panic(err)
	}
	fmt.Println("Success to Migrate")
}
