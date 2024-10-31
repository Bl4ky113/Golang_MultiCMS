package cms_templates

import (
    "fmt"
	"path/filepath"

	"github.com/gin-contrib/multitemplate"
)

var tmplBaseSufixs = []string{"header", "sidebar", "footer"}

func LoadTemplatesToRenderer (htmlRenderer *multitemplate.Renderer, tmplsPath, tmplsSubPath, tmplsPrefix string) {
	layouts, err := filepath.Glob(tmplsPath + tmplsPrefix + "_*.tmpl")
	if err != nil {
        // TODO: Better log
		panic(err.Error())
	}

	templates, err := filepath.Glob(tmplsPath + tmplsSubPath + tmplsPrefix + "_*.tmpl")
	if err != nil {
        // TODO: Better log
		panic(err.Error())
	}

	for _, template := range templates {
		layoutCopy := make([]string, len(layouts))
		copy(layoutCopy, layouts)

        fmt.Printf("%v\n", template)

		layoutCopy = append(layoutCopy, template)
		(*htmlRenderer).AddFromFiles(filepath.Base(template), layoutCopy...)
	}
}
