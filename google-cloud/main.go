package main

import (
	"net/http"
	"restapi/app/service"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
		engine := gin.Default()
    engine.GET("/", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "message": "hello world",
        })
        })

    engine.GET("/test", func(c *gin.Context) {
        userName := service.Regist()
        c.JSON(http.StatusOK, gin.H{
            "message": userName,
        })
        })

    engine.Run()
}
