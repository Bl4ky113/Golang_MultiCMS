package cms_templates

import (
	"strings"

	"github.com/gin-gonic/gin"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

const (
    adminHomePath = "admin"
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
    stringCases := cases.Title(language.Und)

    for _, str := range pathRawArray {
        if strings.Compare(str, adminHomePath) == 0 {
            pathArray = append(pathArray, "Home")

            continue
        }

        if len(str) <= 0 { continue }

        pathArray = append(pathArray, stringCases.String(str))
    }

    return pathArray
}
