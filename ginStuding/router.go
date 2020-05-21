package main

/**
* @Author: 胡大海
* @Date: 2020-05-09 14:21
* A programmer who subconsciously views himself as an artist will enjoy what he does and will do it better ​
 */

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)


func main() {
	r := gin.Default()
	fmt.Println("r.basepath is ", r.BasePath())
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ping",
		})
	})
	//r.GET("/ping/pong", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"message": "pong",
	//	})
	//})
	//r.GET("/user/*action", func(c *gin.Context) {
	//	action := c.Param("action")
	//	message := "action is " + action
	//	c.String(http.StatusOK, message)
	//})
	r.GET("/user/:put", func(c *gin.Context) {
		action := c.Param("put")
		message := "action is " + action
		c.String(http.StatusOK, message)
	})
	//r.GET("/user/:get", func(c *gin.Context) {
	//	action := c.Param("put")
	//	message := "action is " + action
	//	c.String(http.StatusOK, message)
	//})
	r.GET("/user/:put/string", func(c *gin.Context) {
		action := c.Param("put")
		message := "action is " + action
		c.String(http.StatusOK, message)
	})
	r.GET("/ping/put", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	fmt.Println("all routes is ", r.Routes())
	gin.SetMode(gin.DebugMode)
	_ = r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
