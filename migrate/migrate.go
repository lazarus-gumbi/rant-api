package main

import (
	"github.com/lazarus-gumbi/sample_api/initializers"
	"github.com/lazarus-gumbi/sample_api/models"
)

func init() {
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
