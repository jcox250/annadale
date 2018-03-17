package repository

type BaseRepo struct {
	dbHandler DBHandler
}

func NewBaseRepo(dbHandler DBHandler) *BaseRepo {
	return &BaseRepo{
		dbHandler: dbHandler,
	}
}

func (b *BaseRepo) GetLatestResults() ([]string, error) {
	result := ""
	results := []string{}

	rows, err := b.dbHandler.Query("call spBaseGetLatestResults()")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err := rows.Scan(
			&result,
		)
		if err != nil {
			return nil, err
		}
		results = append(results, result)
	}
	return results, nil
}
