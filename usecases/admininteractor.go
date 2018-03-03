package usecases

import (
	"time"

	"github.com/jcox250/annadale/domain"
)

type AdminRepo interface {
	GetNews()
	GetPages()
}

type AdminInteractor struct {
	Repo AdminRepo
}

func NewAdminInteractor(repo AdminRepo) *AdminInteractor {
	return &AdminInteractor{
		Repo: repo,
	}
}

func (a AdminInteractor) GetNews() ([]domain.Post, error) {
	news := []domain.Post{
		domain.Post{
			ID:          "1",
			Title:       "News1",
			Content:     "this is some content",
			AddedBy:     "1",
			AddedDate:   time.Now(),
			UpdatedBy:   "1",
			UpdatedDate: time.Now(),
			Archive:     false,
		},

		domain.Post{
			ID:          "2",
			Title:       "News 2",
			Content:     "this is some content",
			AddedBy:     "1",
			AddedDate:   time.Now(),
			UpdatedBy:   "1",
			UpdatedDate: time.Now(),
			Archive:     false,
		},
	}
	return news, nil
}

func (a AdminInteractor) GetPages() ([]domain.Post, error) {
	pages := []domain.Post{
		domain.Post{
			ID:          "1",
			Title:       "Page 1",
			Content:     "this is some content",
			AddedBy:     "1",
			AddedDate:   time.Now(),
			UpdatedBy:   "1",
			UpdatedDate: time.Now(),
			Archive:     false,
		},

		domain.Post{
			ID:          "2",
			Title:       "Page 2",
			Content:     "this is some content",
			AddedBy:     "1",
			AddedDate:   time.Now(),
			UpdatedBy:   "1",
			UpdatedDate: time.Now(),
			Archive:     false,
		},
	}
	return pages, nil
}
