package service

import (
	"fmt"
	"html/template"
)

const (
	baseTmpl = "infrastructure/view/base.html"
	postTmpl = "infrastructure/view/post.html"
)

const (
	postPage = iota
)

var templates = map[int]*template.Template{
	postPage: template.Must(generateTemplate(baseTmpl, postTmpl)),
}

func generateTemplate(tmpls ...string) (*template.Template, error) {
	t, err := template.ParseFiles(tmpls...)
	if err != nil {
		return nil, fmt.Errorf("couldn't parse %s template: %s", tmpls, err)
	}
	return t, nil
}
