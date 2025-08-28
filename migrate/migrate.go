package main

import (
	"github.com/korshunvladislav/testTaskEM/initializers"
	"github.com/korshunvladislav/testTaskEM/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Subscription{})
}
