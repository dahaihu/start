package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Student struct {
	Name []string `form:"name"`
	Age  *int      `form:"age"`
}

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		var student *Student = &Student{}
		if err := c.Bind(student); err != nil {
			c.JSON(http.StatusInternalServerError, err)
		} else {
			fmt.Printf("Bind result is %v\n", student)
			c.JSON(http.StatusOK, student)
		}
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
