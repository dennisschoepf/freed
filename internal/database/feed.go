package database

type FeedType string

const (
	RSS FeedType = "RSS"
)

type Feed struct {
	Name     string
	Url      string
	FeedType FeedType
}

func (f Feed) Insert() (int64, error) {
	result, err := db.Exec("INSERT INTO feed (name, url, type) VALUES (?,?,?)", f.Name, f.Url, f.FeedType)

	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return 0, err
	}

	return id, nil
}
