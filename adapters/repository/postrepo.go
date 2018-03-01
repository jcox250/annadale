package repository

type PostRepo struct {
	dbHandler DBHandler
}

func NewPostRepo(dbHandler DBHandler) *PostRepo {
	return &PostRepo{
		dbHandler: dbHandler,
	}
}

func (p *PostRepo) Get(id string) {
	p.dbHandler.Execute("call spPostGetPost(?)", id)
}

func (p *PostRepo) GetAll() {
	p.dbHandler.Query("call spPostGetAll()")
}
