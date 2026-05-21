package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hola, Mundo!",
		})
	})

	router.GET("/movies", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Lista de películas",
		})
	})

	if err := router.Run(":8080"); err != nil {
		fmt.Printf("Error al iniciar el servidor: %v\n", err)
		panic(err)
	}

}
