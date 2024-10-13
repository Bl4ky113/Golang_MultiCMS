package main

import (
    pingRoutes "multi_cms/ping"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	rtr := gin.Default()

    // Ping Group
    pingRoutes.AddRootRoutes(rtr)

	return rtr
}

func main() {
	rtr := setupRouter()
    
	rtr.Run("localhost:8000")
}
