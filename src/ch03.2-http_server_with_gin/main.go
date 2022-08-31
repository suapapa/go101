package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/hello", helloHandler)
	// r.Run(":8080")
	if err := r.RunTLS(":443", "server.crt", "server.key"); err != nil {
		panic(err)
	}
}

func helloHandler(c *gin.Context) {
	res := map[string]any{
		"say": "🍸 Cheers",
		"to":  "Gophers",
	}
	if name := c.Query("name"); name != "" {
		res["to"] = name
	}
	c.JSON(http.StatusOK, res)
}
