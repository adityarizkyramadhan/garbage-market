package main

import (
	"github.com/adityarizkyramadhan/garbage-market/infrastructure/database_connection"
	"github.com/adityarizkyramadhan/garbage-market/infrastructure/database_driver"
	"github.com/adityarizkyramadhan/garbage-market/middleware"
	"github.com/adityarizkyramadhan/garbage-market/route"
	"github.com/gin-gonic/gin"
)

func main() {
	// Your code here...
	r := gin.Default()
	r.Use(middleware.CorsMiddleware())
	envDb, err := database_driver.ReadEnvSupabase()
	if err != nil {
		panic(err)
	}
	db, err := database_connection.MakeConnection(envDb)
	if err != nil {
		panic(err)
	}
	route.InitRoute(r, db)
	if err := r.Run(); err != nil {
		panic(err)
	}
}
