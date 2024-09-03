package database

type FeedType string

type Feed struct {
	Name string
	Url  string
}

func (f Feed) Insert() (int64, error) {
	result, err := db.Exec("INSERT INTO feed (name, url) VALUES (?,?)", f.Name, f.Url)

	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return 0, err
	}

	return id, nil
}
