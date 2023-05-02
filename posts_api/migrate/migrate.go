package migrateimport

import (
	"fmt"
	"log"
	"posts_api/initializers"
	"posts_api/models"
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("Could not load env", err)
	}

	initializers.ConnectDB(&config)
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
	fmt.Println("Migration complete")
}
