package ping

import (
    "net/http"

    "github.com/gin-gonic/gin"

    shared_types "multi_cms/shared/cms_types"
)

func ping (cxt *gin.Context) {
    cxt.String(http.StatusOK, "pong")
}

func AddRootRoutes [Groupable shared_types.RouterGroupable](rtr Groupable) {
    rootGrp := rtr.Group("/root")

    rootGrp.GET("/ping", ping)
}
