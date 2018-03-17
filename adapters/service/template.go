package service

import (
	"fmt"
	"html/template"
	"io"

	"github.com/jcox250/annadale/domain"
	"github.com/jcox250/annadale/usecases"
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

type BaseTemplate struct {
	Template *template.Template
	Data     domain.Base
}

func newBaseTemplate(tmpls ...string) BaseTemplate {
	tmpls = append(tmpls, baseTmpl)
	return BaseTemplate{
		Template: template.Must(generateTemplate(tmpls...)),
	}
}

func (b BaseTemplate) ExecuteTemplate(wr io.Writer, name string, data interface{}) error {
	var interactor usecases.BaseInteractor
	var err error
	b.Data.BaseData.LatestResults, err = interactor.GetLatestResults()
	if err != nil {
		return err
	}
	b.Data.BaseData.NextFixtures, err = interactor.GetFixtures()
	if err != nil {
		return err
	}
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
