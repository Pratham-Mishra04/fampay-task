package initializers

import (
	"fmt"
)

func AutoMigrate() {
	fmt.Println("\nStarting Migrations...")
	DB.AutoMigrate()
	fmt.Println("Migrations Finished!")
}
