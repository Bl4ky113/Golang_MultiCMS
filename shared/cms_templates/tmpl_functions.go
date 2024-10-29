package cms_templates

import (
    "html/template"
)

// Not beeing used right now on the templates, 
// but might be use full in the future
func AddTemplateFunctions (tmpl *template.Template) {
    tmplFunctionsMap := template.FuncMap{
        "lststr": tmplLastStr,
        "sub": tmplSubtractValue,
    }

    tmpl.Funcs(tmplFunctionsMap)
}

func tmplLastStr (arr []string) string {
    return arr[len(arr) - 1] 
}

func tmplSubtractValue (a int, b int) int {
    return a - b
}
