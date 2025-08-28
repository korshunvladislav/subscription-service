package main

import (
	"github.com/gin-gonic/gin"
	"github.com/korshunvladislav/testTaskEM/controllers"
	"github.com/korshunvladislav/testTaskEM/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	r.POST("/subscriptions", controllers.SubscriptionsCreate)
	r.GET("/subscriptions/:id", controllers.SubscriptionsShow)
	r.PUT("/subscriptions/:id", controllers.SubscriptionsUpdate)
	r.DELETE("/subscriptions/:id", controllers.SubscriptionsDelete)
	r.GET("/subscriptions", controllers.SubscriptionsIndex)

	r.GET("/subscriptions/summary", controllers.SubscriptionsSummary)

	r.Run()
}
