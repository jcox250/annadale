package service

import (
	"fmt"
	"html/template"
)

const (
	headerTmpl = "infrastructure/view/header.html"
	footerTmpl = "infrastructure/view/footer.html"
	postTmpl   = "infrastructure/view/post.html"
)

const (
	postPage = iota
)

var templates = map[int]*template.Template{
	postPage: template.Must(generateTemplate(postTmpl)),
}

func generateTemplate(tmpls ...string) (*template.Template, error) {
	// Always include header and footer templates
	tmpls = append(tmpls, headerTmpl)
	tmpls = append(tmpls, footerTmpl)
	t, err := template.ParseFiles(tmpls...)
	if err != nil {
		return nil, fmt.Errorf("couldn't parse %s template: %s", tmpls, err)
	}
	return t, nil
}
