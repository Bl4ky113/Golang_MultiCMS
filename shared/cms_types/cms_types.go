package cms_types

import (
    "github.com/gin-gonic/gin"
)

type RouterGroupable interface {
    *gin.Engine | *gin.RouterGroup;
    Group(string, ...gin.HandlerFunc) *gin.RouterGroup;
}
