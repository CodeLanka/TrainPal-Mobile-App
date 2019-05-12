package main

import (
	routes "CodeLanka/TrainPal-Mobile-App/go-app/routes"
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {

	r := gin.Default()
	routes.BindRoutes(r)
	return r
}

func main() {
	r := setupRouter()

	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
