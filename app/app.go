package app

import (
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	mapEndpoints()

	// Start App at localhost:8000
	router.Run(":8000")
}
