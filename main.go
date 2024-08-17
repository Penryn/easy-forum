package main

import (
	"easy-forum/internal/pkg/database"
	"easy-forum/internal/router"
	"easy-forum/internal/service"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	db := database.Init()
	service.ServiceInit(db)
	r := gin.Default()
	router.Init(r)
	err := r.Run(":8081")
	if err != nil {
		log.Fatal(err)
	}
}
