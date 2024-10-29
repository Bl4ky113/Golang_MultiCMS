package cms_templates

import (
	"fmt"
	"html/template"
)

var tmplBaseSufixs = []string{"header", "sidebar", "footer"}

func CreateTemplateFromBase (tmplGroup *template.Template, tmplAddonName, tmplPrefix, addonName string) (*template.Template, error) {
    tmplBase, err := tmplGroup.Lookup(tmplPrefix + "_base.tmpl").Clone()
    if tmplBase == nil || err != nil {
        // TODO: Fatal-Error msg searching and copying for template
    }
   
    tmplAddon := tmplGroup.Lookup(tmplPrefix + "_" + tmplAddonName)
    if tmplAddon == nil {
        // TODO: Fatal-Error msg searching for template
    }

    _, err = tmplBase.AddParseTree(tmplPrefix + "_main_content", tmplAddon.Tree)
    if err != nil {
        // TODO: Fatal-Error msg
    }

    return tmplBase, nil
}

func AddBasesToTemplate (tmpl, tmplGroup *template.Template, tmplPrefix string) (error) {
    for _, sufix := range tmplBaseSufixs {
        tmplElement, err := tmplGroup.Lookup(tmplPrefix + "_" + sufix + ".tmpl").Clone()
        if tmplElement == nil || err != nil {
            // TODO: Fatal-Error msg
        }
        
        foo, err := tmpl.AddParseTree(
            tmplPrefix + "_" + sufix,
            tmplElement.Tree,
        )

        fmt.Println(foo.Name())
        for _, bar := range foo.Templates() {
            fmt.Printf("\t\t-%v\n", bar.Name())
        }
    }

    return nil
}
