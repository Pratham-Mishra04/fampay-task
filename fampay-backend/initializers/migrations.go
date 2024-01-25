package initializers

import (
	"fmt"

	"github.com/Pratham-Mishra04/fampay/fampay-backend/models"
)

func AutoMigrate() {
	fmt.Println("\nStarting Migrations...")
	DB.AutoMigrate(&models.Video{})
	fmt.Println("Migrations Finished!")
}
