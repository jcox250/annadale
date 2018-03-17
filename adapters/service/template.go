package service

import (
	"fmt"
	"html/template"
	"io"
)

const (
	baseTmpl     = "infrastructure/view/base.html"
	postTmpl     = "infrastructure/view/post.html"
	homeTmpl     = "infrastructure/view/home.html"
	adminTmpl    = "infrastructure/view/admin.html"
	loginTmpl    = "infrastructure/view/login.html"
	editPostTmpl = "infrastructure/view/edit-post.html"
	notFoundTmpl = "infrastructure/view/not-found.html"
)

const (
	postPage = iota
	homePage
	adminPage
	loginPage
	editPostPage
	notFoundPage
)

type Base struct {
	BaseData interface{}
	PageData interface{}
}

type BaseTemplate struct {
	Template *template.Template
	Data     Base
}

func newBaseTemplate(tmpls ...string) BaseTemplate {
	tmpls = append(tmpls, baseTmpl)
	b := BaseTemplate{
		Template: template.Must(generateTemplate(tmpls...)),
		Data:     Base{BaseData: "hello world"},
	}
	return b
}

func (b BaseTemplate) ExecuteTemplate(wr io.Writer, name string, data interface{}) error {
	b.Data.PageData = data
	return b.Template.ExecuteTemplate(wr, name, b.Data)
}

var templates = map[int]BaseTemplate{
	postPage:     newBaseTemplate(postTmpl),
	homePage:     newBaseTemplate(homeTmpl),
	adminPage:    newBaseTemplate(adminTmpl),
	loginPage:    newBaseTemplate(loginTmpl),
	editPostPage: newBaseTemplate(editPostTmpl),
	notFoundPage: newBaseTemplate(notFoundTmpl),
}

func generateTemplate(tmpls ...string) (*template.Template, error) {
	t, err := template.ParseFiles(tmpls...)
	if err != nil {
		return nil, fmt.Errorf("couldn't parse %s template: %s", tmpls, err)
	}
	return t, nil
}
