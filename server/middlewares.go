package server

import "github.com/gin-gonic/gin"

type MiddlewareFunc func(c *gin.Context)

func GetMiddlewares() []MiddlewareFunc {
	return []MiddlewareFunc{
		Logging,
	}
}

func Logging(c *gin.Context) {
	//TODO: Add logger
}
