package cms_templates

import (
    "html/template"
)

// Not beeing used right now on the templates, 
// but might be use full in the future
func GetTemplateFunctions () (template.FuncMap) {
    tmplFunctionsMap := template.FuncMap{
        "lststr": tmplLastStr,
        "sub": tmplSubtractValue,
    }

    return tmplFunctionsMap
}

func tmplLastStr (arr []string) string {
    return arr[len(arr) - 1] 
}

func tmplSubtractValue (a int, b int) int {
    return a - b
}
