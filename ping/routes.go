package ping

import (
    "net/http"

    "github.com/gin-gonic/gin"

    sharedTypes "multi_cms/shared/cms_types"
)

func ping (cxt *gin.Context) {
    cxt.String(http.StatusOK, "pong")
}

func AddRootRoutes [Groupable sharedTypes.RouterGroupable](rtr Groupable) {
    rootGrp := rtr.Group("/root")

    rootGrp.GET("/ping", ping)
}
