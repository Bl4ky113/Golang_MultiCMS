package cms_templates

import (
    "strings"

    "github.com/gin-gonic/gin"
)

const (
    ADMIN_HOME_PATH = "admin"
)

func GenerateViewContext (ctx *gin.Context, viewCtx *gin.H) error {
    
    // Get View Data
    (*viewCtx)["view"] = gin.H{
        "pathArray": generatePathArray(ctx.FullPath()),
        "usrName": "AdminUSR",
        "usrRole": "Admin",
        "usrIcon": "U",
    }

    return nil
}

func generatePathArray (pathRawString string) []string {
    pathArray := make([]string, 0, 0)
    pathRawArray := strings.Split(pathRawString, "/")

    for _, str := range pathRawArray {
        if strings.Compare(str, ADMIN_HOME_PATH) == 0 {
            pathArray = append(pathArray, "Home")

            continue
        }

        if len(str) <= 0 { continue }

        pathArray = append(pathArray, str)
    }

    return pathArray
}
