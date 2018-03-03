package usecases

import (
	"time"

	"github.com/jcox250/annadale/domain"
)

type PostRepo interface {
	GetPost(string)
	GetAllPosts()
	AddPost(domain.Post) (int64, error)
	UpdatePost(domain.Post) (bool, error)
	GetNews() ([]domain.Post, error)
	GetPages() ([]domain.Post, error)
}

type PostInteractor struct {
	postRepo PostRepo
}

func NewPostInteractor(postRepo PostRepo) PostInteractor {
	return PostInteractor{
		postRepo: postRepo,
	}
}

func (p PostInteractor) GetPost(id string) domain.Post {
	p.postRepo.GetPost(id)

	foo := []domain.Post{
		domain.Post{
			ID:          "1",
			Title:       "Title",
			Content:     "this is some content",
			AddedBy:     "1",
			AddedDate:   time.Now(),
			UpdatedBy:   "1",
			UpdatedDate: time.Now(),
			Archive:     false,
		},

		domain.Post{
			ID:          "2",
			Title:       "Post 2",
			Content:     "this is some content",
			AddedBy:     "1",
			AddedDate:   time.Now(),
			UpdatedBy:   "1",
			UpdatedDate: time.Now(),
			Archive:     false,
		},
	}

	for _, post := range foo {
		if post.ID == id {
			return post
		}
	}
	return foo[1]
}

func (p PostInteractor) GetAllPosts() {
	p.postRepo.GetAllPosts()
}

func (a PostInteractor) GetNews() ([]domain.Post, error) {
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

func (a PostInteractor) GetPages() ([]domain.Post, error) {
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

func (p PostInteractor) AddPost(post domain.Post) (int64, error) {
	return p.postRepo.AddPost(post)
}

func (p PostInteractor) EditPost(post domain.Post) (bool, error) {
	return p.postRepo.UpdatePost(post)
}
