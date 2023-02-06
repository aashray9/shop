package main

import (
	"net/http"
	"os"

	"lms/common"
	"lms/controllers"
	"lms/database"
	"lms/middleware"

	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()
	app.Use(gin.Recovery())

	godotenv.Load()
	common.Loggers()

	app.Use(middleware.RequestLogger)
	db := database.Connection()

	SqlDB, _ := db.DB()
	defer SqlDB.Close()
	gin.SetMode(os.Getenv("GIN_MODE"))

	app.Use(middleware.JwtTokenVerify())

	apiRoutes := app.Group("/api")
	controllers.BaseRoutes(apiRoutes)

	controllers.LeadRoutes(apiRoutes.Group("/lead_managements"))

	app.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusBadRequest, gin.H{
			"success":  false,
			"messages": "No route find!!!",
		})
	})

	app.Run(os.Getenv("APP_PORT"))
}
