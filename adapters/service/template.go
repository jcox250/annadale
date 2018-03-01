package service

import (
	"fmt"
	"text/template"
)

const (
	postTemplate = "post.html"
)

func generateTemplate(tmpl string) (*template.Template, error) {
	file := "infrastructure/view/" + tmpl
	t, err := template.ParseFiles(file)
	if err != nil {
		return nil, fmt.Errorf("couldn't parse %s template: %s", tmpl, err)
	}
	return t, nil
}
