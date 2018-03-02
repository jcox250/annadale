package usecases

import (
	"time"

	"github.com/jcox250/annadale/domain"
)

type HomeRepo interface {
	GetTopNews() ([]domain.Post, error)
}

type HomeInteractor struct {
	HomeRepo HomeRepo
}

func NewHomeInteractor(repo HomeRepo) HomeInteractor {
	return HomeInteractor{
		HomeRepo: repo,
	}
}

func (h HomeInteractor) GetTopNews() ([]domain.Post, error) {
	topNews := []domain.Post{
		domain.Post{
			ID:          "1",
			Title:       "Won the cup",
			Content:     "this is some content",
			AddedBy:     "1",
			AddedDate:   time.Now(),
			UpdatedBy:   "1",
			UpdatedDate: time.Now(),
			Archive:     false,
		},

		domain.Post{
			ID:          "2",
			Title:       "Quiz Night",
			Content:     "this is some content",
			AddedBy:     "1",
			AddedDate:   time.Now(),
			UpdatedBy:   "1",
			UpdatedDate: time.Now(),
			Archive:     false,
		},
	}
	return topNews, nil
}
