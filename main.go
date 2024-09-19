package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"fileserver/config"
	"fileserver/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// constants
const (
	PORT = 4999
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	app := gin.New()

	// max memory setup - Set 10GB max upload size
	app.MaxMultipartMemory = 10000000000

	//Cors
	app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowHeaders:     []string{"R-token", "Origin", "Content-Length", "Content-Type", "Authorization", "X-Requested-With", "Accept", "Access-Control-Allow-Origin"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
		AllowWebSockets:  true,
	}))

	app.POST("/upload", controllers.UploadFile)

	// Use the fileserver package to serve static files
	app.Static(config.FILES_ACCESS_PREFIX_URL, config.FILES_STATIC_DIR)

	app.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"message": "Page not found"})
	})

	fmt.Println("Server is running on port http://localhost:" + strconv.Itoa(PORT))
	app.Run(fmt.Sprintf(":%d", PORT))
}
