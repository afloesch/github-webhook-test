package main

import (
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func main() {
	// Creates a router without any middleware by default
	r := gin.New()

	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	r.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())

	// Per route middleware, you can add as many as you desire.
	r.POST("/", func(c *gin.Context) {
		body := c.Request.Body
		x, _ := ioutil.ReadAll(body)

		fmt.Printf("%s \n", string(x))
	})

	// Listen and serve on 0.0.0.0:8080
	r.Run(":8080")
}
