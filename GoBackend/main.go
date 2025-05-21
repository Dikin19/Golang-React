package main

import (
	"MuhamadSodikin/backend-api/config"
	"MuhamadSodikin/backend-api/database"
	"MuhamadSodikin/backend-api/routes"
)

func main() {

	//load config .env
	config.LoadEnv()

	//inisialisasi database
	database.InitDB()

	//inisialiasai Gin
	r := routes.SetupRouter()

	// //membuat route dengan method GET
	// router.GET("/", func(c *gin.Context) {

	// 	//return response JSON
	// 	c.JSON(200, gin.H{
	// 		"message": "Hello World!",
	// 	})
	// })

	//mulai server
	r.Run(":" + config.GetEnv("APP_PORT", "3001"))
}