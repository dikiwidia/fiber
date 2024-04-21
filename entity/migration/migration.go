package migration

import (
	"fmt"

	"github.com/dikiwidia/fiber/config"
	"github.com/dikiwidia/fiber/entity"
)

func RunMigration() {
	err := config.DB.AutoMigrate(&entity.User{})
	if err != nil {
		panic(err)
	}
	fmt.Println("Success to Migrate")
}
