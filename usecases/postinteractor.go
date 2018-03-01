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

	return domain.Post{
		ID:          "123",
		Title:       "Title",
		Content:     "this is some content",
		AddedBy:     "1",
		AddedDate:   time.Now(),
		UpdatedBy:   "1",
		UpdatedDate: time.Now(),
		Archive:     false,
	}
}

func (p PostInteractor) GetAllPosts() {
	p.postRepo.GetAll()
}
