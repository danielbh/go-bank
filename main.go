package main

import (
	"github.com/gin-gonic/gin"

	"github.com/danielbh/go-bank/middleware"
	"github.com/danielbh/go-bank/transactions"
)

func main() {
	// := same as var <T> = gin.Default
	// var x int = 2
	// var x = 2
	// x := 2
	router := gin.Default()
	router.Use(middleware.RequestIDMiddleware())
	v1 := router.Group("/api")
	transactions.Register(v1.Group("/transactions"))

	router.Run("localhost:8080")
}
