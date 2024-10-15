package main

import (
    pingRoutes "multi_cms/ping"
    adminRoutes "multi_cms/admin"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	rtr := gin.Default()

    // Ping Group
    pingRoutes.AddRootRoutes(rtr)

    // Admin Group
    adminRoutes.AddAdminTemplates(rtr)
    adminRoutes.AddAdminRoutes(rtr)

	return rtr
}

func main() {
	rtr := setupRouter()
    
	rtr.Run("localhost:8000")
}
