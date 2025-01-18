package products

import (
    "net/http"

    "github.com/gin-gonic/gin"

    shrTemplates "multi_cms/shared/cms_templates"
    shrLogger "multi_cms/shared/cms_logging"
)

func ViewProducts (ctx *gin.Context) {
    viewCtx := make(gin.H)

    err := shrTemplates.GenerateViewContext(ctx, &viewCtx)
    if err != nil {
        // TODO: ADD cms_logging error log
    }

    shrLogger.InfoLog(shrLogger.ERR_CODE_ADMIN_PRODUCTS, "TEST: ACCESSING Products")
    ctx.HTML(http.StatusOK, "admin_products.tmpl", viewCtx)
}
