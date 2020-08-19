package main

import (
	"github.com/gin-gonic/gin"
)

// GetDummyEndpoint Test
func GetDummyEndpoint(c *gin.Context) {
	resp := map[string]string{"hello": "world"}
	c.JSON(200, resp)
}

// DummyMiddleware test
func DummyMiddleware() gin.HandlerFunc {
	// Do some initialization logic here
	// Foo()
	return func(c *gin.Context) {
		c.Next()
	}
}
