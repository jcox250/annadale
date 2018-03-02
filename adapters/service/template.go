package service

import (
	"fmt"
	"html/template"
)

const (
	baseTmpl = "infrastructure/view/base.html"
	postTmpl = "infrastructure/view/post.html"
	homeTmpl = "infrastructure/view/home.html"
)

const (
	postPage = iota
	homePage
)

var templates = map[int]*template.Template{
	postPage: template.Must(generateTemplate(baseTmpl, postTmpl)),
	homePage: template.Must(generateTemplate(baseTmpl, homeTmpl)),
}

func generateTemplate(tmpls ...string) (*template.Template, error) {
	t, err := template.ParseFiles(tmpls...)
	if err != nil {
		return nil, fmt.Errorf("couldn't parse %s template: %s", tmpls, err)
	}
	return t, nil
}
