package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/multitemplate"

	shrTemplates "multi_cms/shared/cms_templates"
	shrTypes "multi_cms/shared/cms_types"

	"multi_cms/admin/home"
	"multi_cms/admin/products"
)

const (
    templatesPath = "admin/resources/html/"
    templatesPrefix = "admin"
    cssPath = "admin/resources/css/"
    jsPath = "admin/resources/js/"
)

func AddAdminRoutes [Groupable shrTypes.RouterGroupable](rtr Groupable) {
    adminGrp := rtr.Group("/admin")

    adminGrp.GET("/", home.ViewHome)
    adminGrp.GET("/products/", products.ViewProducts)
}

func AddAdminTemplates (rtr *gin.Engine, htmlRenderer *multitemplate.Renderer) {
    shrTemplates.LoadTemplatesToRenderer(htmlRenderer, templatesPath, "home/", templatesPrefix)
    shrTemplates.LoadTemplatesToRenderer(htmlRenderer, templatesPath, "products/", templatesPrefix)

    // Add static files
    rtr.Static("/admin/css/", cssPath)
    rtr.Static("/admin/js/", jsPath)
}
