package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/multitemplate"

	adminRoutes "multi_cms/admin"
	pingRoutes "multi_cms/ping"

	shrTemplates "multi_cms/shared/cms_templates"
)

func setupRouter() *gin.Engine {
	rtr:= gin.Default()

    // Config Gin
    gin.SetMode(gin.DebugMode)

    // Config and Add Templates
    setupTemplates(rtr) 

    // Setup Routes and RouteGroups
    setupRoutes(rtr)

	return rtr
}

func setupTemplates (rtr *gin.Engine) {
    htmlRender := multitemplate.NewRenderer()

    // Template Custom Functions
    rtr.SetFuncMap(shrTemplates.GetTemplateFunctions())

    // Add Apps Routes
    adminRoutes.AddAdminTemplates(rtr, &htmlRender)

    rtr.HTMLRender = htmlRender
}

func setupRoutes (rtr *gin.Engine) {
    // Ping Group
    pingRoutes.AddRootRoutes(rtr)

    // Admin Group
    adminRoutes.AddAdminRoutes(rtr)
}

func main() {
	rtr := setupRouter()
    
	rtr.Run("localhost:8000")
}
