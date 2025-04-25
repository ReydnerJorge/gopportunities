package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// Inicializa a Router utilizando as configs Default do gin
	r *gin.Engine := gin.Default()
	// Definindo uma rota
	r.GET(relativePath: "/ping", handlers ... : func(c *gin.Context){
		c.JSON(code: 200, obj: gin.H {
			"message": "pong",
		})
	})
	// Estamos rodando a nossa api
	r.Run() //listen and server on 0.0.0.0:8080
}