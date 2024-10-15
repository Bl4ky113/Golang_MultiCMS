package admin

import (
    "github.com/gin-gonic/gin"

    shrTypes "multi_cms/shared/cms_types"

    "multi_cms/admin/home"
)

const (
    templatesPath = "admin/resources/html/*.tmpl"
    cssPath = "admin/resources/css/"
    jsPath = "admin/resources/js/"
)

func AddAdminRoutes [Groupable shrTypes.RouterGroupable](rtr Groupable) {
    adminGrp := rtr.Group("/admin")

    adminGrp.GET("/", home.ViewHome)
}

func AddAdminTemplates (rtr *gin.Engine) {
    // We'll see later if this bites us in the ass
    rtr.LoadHTMLGlob(templatesPath)
    rtr.Static("/admin/css/", cssPath)
    rtr.Static("/admin/js/", jsPath)
}
