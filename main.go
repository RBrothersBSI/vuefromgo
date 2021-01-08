package main

import (
	"fmt"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	// Serve static files from the frontend/dist directory.
	r := gin.Default()

	// Serve frontend static files
	r.Use(static.Serve("/", static.LocalFile("./frontend/dist", true)))

	// Setup route group for the API
	api := r.Group("/api")
	{
		api.GET("/hello",homePageHandler)
	}


	// Start the server.
	fmt.Println("Server listening on port 3000")
	r.Run(":3000")

}

func homePageHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func checkError(err error) {
	if err != nil {
		log.Panic(err)
	}
}
