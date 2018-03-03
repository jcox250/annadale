package service

import (
	"fmt"
	"html/template"
)

const (
	baseTmpl     = "infrastructure/view/base.html"
	postTmpl     = "infrastructure/view/post.html"
	homeTmpl     = "infrastructure/view/home.html"
	adminTmpl    = "infrastructure/view/admin.html"
	loginTmpl    = "infrastructure/view/login.html"
	addPostTmpl  = "infrastructure/view/add-post.html"
	notFoundTmpl = "infrastructure/view/not-found.html"
)

const (
	postPage = iota
	homePage
	adminPage
	loginPage
	addPostPage
	notFoundPage
)

var templates = map[int]*template.Template{
	postPage:     template.Must(generateTemplate(baseTmpl, postTmpl)),
	homePage:     template.Must(generateTemplate(baseTmpl, homeTmpl)),
	adminPage:    template.Must(generateTemplate(baseTmpl, adminTmpl)),
	loginPage:    template.Must(generateTemplate(baseTmpl, loginTmpl)),
	addPostPage:  template.Must(generateTemplate(baseTmpl, addPostTmpl)),
	notFoundPage: template.Must(generateTemplate(baseTmpl, notFoundTmpl)),
}

func generateTemplate(tmpls ...string) (*template.Template, error) {
	t, err := template.ParseFiles(tmpls...)
	if err != nil {
		return nil, fmt.Errorf("couldn't parse %s template: %s", tmpls, err)
	}
	return t, nil
}
