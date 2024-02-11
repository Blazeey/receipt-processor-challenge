package server

import (
	"net/http"
	"receipt-processor-challenge/models"
	routes "receipt-processor-challenge/server/routes"
	"receipt-processor-challenge/utils"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// Creating the gin router and initializing routes
func InitializeHttpRoutes() *gin.Engine {
	router := gin.New()
	gin.SetMode(gin.ReleaseMode)
	router.Use(gin.Logger())
	router.Use(gin.CustomRecovery(recoveryHandler()))

	// Register custom validation rule
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("regex", utils.ValidateRegex)
		_ = v.RegisterValidation("time", utils.ValidateTimeOnly)
	}

	baseRouter := router.Group("/")
	routes.InitializeHeartbeatRoute(baseRouter)
	routes.InitializeReceiptRoutes(baseRouter)
	return router
}

func recoveryHandler() gin.RecoveryFunc {
	return func(c *gin.Context, err any) {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Message: "An error occurred when processing the request"})
	}
}
