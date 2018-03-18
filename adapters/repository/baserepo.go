package repository

type BaseRepo struct {
	dbHandler DBHandler
}

func NewBaseRepo(dbHandler DBHandler) *BaseRepo {
	return &BaseRepo{
		dbHandler: dbHandler,
	}
}

func (b *BaseRepo) GetResults() ([]string, error) {
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
	return []string{
		"1s - Instonians",
		"2s - Three Rock",
		"3s - North Down",
		"4s - Civil Service",
	}, nil
}

func (b *BaseRepo) GetFixtures() ([]string, error) {
	fixture := ""
	fixtures := []string{}

	rows, err := b.dbHandler.Query("call spBaseGetFixtures()")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err := rows.Scan(
			&fixture,
		)
		if err != nil {
			return nil, err
		}
		fixtures = append(fixtures, fixture)
	}
	return []string{
		"1s - Banbridge",
		"2s - Antrim",
		"3s - Raphoe",
		"4s - Bangor",
	}, nil
}
