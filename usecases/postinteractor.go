package usecases

import (
	"time"

	"github.com/jcox250/annadale/domain"
)

type PostInteractor struct {
	postRepo Repo
}

func NewPostInteractor(postRepo Repo) PostInteractor {
	return PostInteractor{
		postRepo: postRepo,
	}
}

func (p PostInteractor) GetPost(id string) domain.Post {
	p.postRepo.Get(id)

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
	p.postRepo.GetAll()
}
