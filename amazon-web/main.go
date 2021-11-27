package main

import (
	"fmt"
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

	engine.GET("/text", func(c *gin.Context) {
		id := c.Query("id")
		fmt.Print(id)
		c.JSON(http.StatusOK, gin.H{
			"id": id,
		})
	})

	engine.GET("/text/:id", func(c *gin.Context) {
		id := c.Param("id")
		fmt.Print(id)
		c.JSON(http.StatusOK, gin.H{
			"id": id,
		})
	})

	engine.POST("/test", func(c *gin.Context) {
		name := c.PostForm("name")
		user := service.Regist(name)
		c.JSON(http.StatusOK, gin.H{
			"userId":   user.Id,
			"userName": user.Name,
		})
	})

	engine.Run()
}
