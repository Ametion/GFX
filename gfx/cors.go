package gfx

import (
	"strings"
)

func AllowAllCORS() MiddlewareFunc {
	return func(c *Context) {
		c.Writer.Header().Set("Access-Control-Allow-Methods", "*")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Next()
	}
}

func AllowMethodsCors(methods ...string) MiddlewareFunc {
	allowedMethods := strings.Join(methods, ", ")

	return func(c *Context) {
		c.Writer.Header().Set("Access-Control-Allow-Methods", allowedMethods)
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Next()
	}
}

func AllowOriginCors(origins ...string) MiddlewareFunc {
	allowedOrigins := strings.Join(origins, ", ")

	return func(c *Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", allowedOrigins)
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Next()
	}
}