package admin

import (
    "github.com/gin-gonic/gin"

    shared_types "multi_cms/shared/cms_types"

    "multi_cms/admin/home"
)

const (
    templatesPath = "admin/resources/html/*.tmpl"
    cssPath = "admin/resources/css/*.css"
    jsPath = "admin/resources/js/*.js"
)

func AddAdminRoutes [Groupable shared_types.RouterGroupable](rtr Groupable) {
    adminGrp := rtr.Group("/admin")

    adminGrp.GET("/", home.ViewHome)
}

func AddAdminTemplates (rtr *gin.Engine) {
    rtr.LoadHTMLGlob("admin/resources/html/*")
}
