package admin

import (
	"fmt"
	"html/template"

	"github.com/gin-gonic/gin"

	shrTypes "multi_cms/shared/cms_types"
    shrTemplates "multi_cms/shared/cms_templates"

	"multi_cms/admin/home"
	"multi_cms/admin/products"
)

const (
    templatesPath = "admin/resources/html/"
    cssPath = "admin/resources/css/"
    jsPath = "admin/resources/js/"
)

func AddAdminRoutes [Groupable shrTypes.RouterGroupable](rtr Groupable) {
    adminGrp := rtr.Group("/admin")

    adminGrp.GET("/", home.ViewHome)
    adminGrp.GET("/products/", products.ViewProducts)
}

func AddAdminTemplates (rtr *gin.Engine, htmlTmpl *template.Template) {
    adminTmpls := template.Must(template.New("").ParseGlob(templatesPath + "*.tmpl"))
    fmt.Println("Admin")
    for _, temp := range adminTmpls.Templates() {
        fmt.Printf("\t- %v\n", temp.Name())
    }

    homeTmpl, err := shrTemplates.CreateTemplateFromBase(
        adminTmpls, "home.tmpl", "admin", "admin_home",
    )
    if err != nil {
        // TODO: Handle Fatal-Error 
    }
    
    shrTemplates.AddBasesToTemplate(homeTmpl, adminTmpls, "admin")
    shrTemplates.AddBasesToTemplate(htmlTmpl, adminTmpls, "admin")

    fmt.Println("Home")
    for _, temp := range homeTmpl.Templates() {
        fmt.Printf("\t- %v\n", temp.Name())
    }

    fmt.Println("html")
    htmlTmpl.AddParseTree("admin_home", homeTmpl.Tree)

    for _, temp := range htmlTmpl.Templates() {
        fmt.Printf("\t- %v\n", temp.Name())
    }

    rtr.Static("/admin/css/", cssPath)
    rtr.Static("/admin/js/", jsPath)
}
