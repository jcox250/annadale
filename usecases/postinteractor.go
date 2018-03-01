package usecases

type PostInteractor struct {
	postRepo Repo
}

func NewPostInteractor(postRepo Repo) PostInteractor {
	return PostInteractor{
		postRepo: postRepo,
	}
}

func (p PostInteractor) GetPost(id string) {
	p.postRepo.Get(id)
}

func (p PostInteractor) GetAllPosts() {
	p.postRepo.GetAll()
}
