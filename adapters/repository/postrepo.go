package repository

import "github.com/jcox250/annadale/domain"

type PostRepo struct {
	dbHandler DBHandler
}

func NewPostRepo(dbHandler DBHandler) *PostRepo {
	return &PostRepo{
		dbHandler: dbHandler,
	}
}

func (p *PostRepo) GetPost(id string) {
	p.dbHandler.Execute("call spPostGetPost(?)", id)
}

func (p *PostRepo) GetAllPosts() {
	p.dbHandler.Query("call spPostGetAll()")
}

func (p *PostRepo) GetNews() ([]domain.Post, error) {
	p.dbHandler.Query("call spPostGetNews()")
	return []domain.Post{}, nil
}

func (p *PostRepo) GetPages() ([]domain.Post, error) {
	p.dbHandler.Query("call spPostGetNews()")
	return []domain.Post{}, nil
}

func (p *PostRepo) AddPost(post domain.Post) (int64, error) {
	rowsAdded, err := p.dbHandler.Execute("call spPostAdd(?,?,?,?,?,?,?,?",
		post.ID,
		post.Title,
		post.Content,
		post.AddedBy,
		post.AddedDate,
		post.UpdatedBy,
		post.UpdatedDate,
		post.Archive,
	)
	return rowsAdded, err
}

func (p *PostRepo) UpdatePost(domain.Post) (bool, error) {
	return false, nil
}
