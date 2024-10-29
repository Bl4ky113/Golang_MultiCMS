package main

import (
	"html/template"

	"github.com/gin-gonic/gin"

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
    htmlTemplates := template.New("")
    shrTemplates.AddTemplateFunctions(htmlTemplates)

    adminRoutes.AddAdminTemplates(rtr, htmlTemplates)

    rtr.SetHTMLTemplate(htmlTemplates)
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
