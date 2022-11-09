package main

import (
	_ "github.com/go-sql-driver/mysql"

	"github.com/gin-gonic/gin"
	"github.com/yunus-mohammed/golang-project/controllers"
	"github.com/yunus-mohammed/golang-project/redisclient"
)

func main() {
	router := gin.Default()
	redisclient.ClientInit()
	router.POST("/test", controllers.RespController)
	router.Run("localhost:8087")
}
