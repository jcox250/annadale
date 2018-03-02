package service

import (
	"fmt"
	"text/template"
)

const (
	headerTmpl = "infrastructure/view/post.html"
	postTmpl   = "infrastructure/view/post.html"
)

func generateTemplate(tmpl ...string) (*template.Template, error) {
	t, err := template.ParseFiles(tmpl...)
	if err != nil {
		return nil, fmt.Errorf("couldn't parse %s template: %s", tmpl, err)
	}
	return t, nil
}
