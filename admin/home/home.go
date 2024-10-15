package home

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

func ViewHome (ctx *gin.Context) {
    ctx.HTML(http.StatusOK, "admin_home.tmpl", gin.H{})
}
