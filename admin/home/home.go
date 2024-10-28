package home

import (
    "net/http"

    "github.com/gin-gonic/gin"

     shrd_templates "multi_cms/shared/cms_templates"
)

func ViewHome (ctx *gin.Context) {
    viewCtx := make(gin.H)

    err := shrd_templates.GenerateViewContext(ctx, &viewCtx)
    if err != nil {
        // TODO: ADD cms_logging error log
    }

    ctx.HTML(http.StatusOK, "admin_home.tmpl", viewCtx)
}
