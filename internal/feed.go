package internal

import (
	"fmt"
	"freed/internal/database"
	"net/url"
	"sync"
	"time"

	"github.com/mmcdole/gofeed"
)

type Feed struct {
	ID           int64
	Name         string
	Url          string
	CreatedAt    *time.Time
	LastSyncedAt *time.Time
	ArticleCount int
}

func AddFeed(feedUrl string, addRecentCount int) (string, int, error) {
	if _, err := url.ParseRequestURI(feedUrl); err != nil {
		return "", 0, fmt.Errorf("The given URL does not seem to be valid: %s", err)
	}

	feed, err := parseByUrl(feedUrl)

	if err != nil {
		return "", 0, err
	}

	f := database.FeedEntity{
		Name: feed.Title,
		Url:  feedUrl,
	}

	feedId, err := f.Insert()
	if err != nil {
		return "", 0, err
	}

	articles := make([]database.ArticleEntity, 0, addRecentCount)

	for i, v := range feed.Items {
		if i > addRecentCount {
			continue
		}

		a := database.ArticleEntity{
			Name:   v.Title,
			Url:    v.Link,
			FeedId: feedId,
		}

		articles = append(articles, a)
	}

	if err := database.InsertMultipleArticles(articles); err != nil {
		return feed.Title, 0, nil
	}

	return feed.Title, len(articles), nil
}

func GetAllFeeds() (*[]Feed, error) {
	feedEntities, err := database.FindAllFeedsWithArticleCount()

	if err != nil {
		return nil, err
	}

	var feeds []Feed

	for _, feedEntity := range *feedEntities {
		feed := &Feed{
			ID:           feedEntity.ID,
			Name:         feedEntity.Name,
			Url:          feedEntity.Url,
			CreatedAt:    feedEntity.CreatedAt,
			LastSyncedAt: feedEntity.LastSyncedAt,
			ArticleCount: feedEntity.ArticleCount,
		}

		feeds = append(feeds, *feed)
	}

	return &feeds, nil
}

func SyncFeeds() error {
	now := time.Now()
	startOfToday := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	feeds, err := database.FindAllFeeds()

	if err != nil {
		return err
	}

	var syncWg sync.WaitGroup
	errChannel := make(chan error, len(*feeds))

	for _, feed := range *feeds {
		syncWg.Add(1)
		go syncFeed(feed, startOfToday, &syncWg, errChannel)
	}

	syncWg.Wait()
	close(errChannel)

	var errors []error

	for err := range errChannel {
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		compositeErrMsg := fmt.Sprintf("Encountered %d errors:\n", len(errors))

		for _, err := range errors {
			compositeErrMsg = compositeErrMsg + " " + fmt.Sprintf("- %v\n", err)
		}

		return fmt.Errorf(compositeErrMsg)
	}

	return nil
}

func syncFeed(
	feed database.FeedEntity,
	syncBefore time.Time,
	wg *sync.WaitGroup,
	errorChannel chan<- error,
) {
	defer wg.Done()

	if feed.LastSyncedAt.After(syncBefore) {
		return
	}

	gFeed, err := parseByUrl(feed.Url)

	if err != nil {
		errorChannel <- err
		return
	}

	var articles []database.ArticleEntity

	for _, item := range gFeed.Items {
		if item.PublishedParsed.Before(*feed.LastSyncedAt) {
			continue
		}

		article := database.ArticleEntity{
			Name:   item.Title,
			Url:    item.Link,
			FeedId: feed.ID,
		}

		articles = append(articles, article)
	}

	if err := database.InsertIgnoreMultipleArticles(articles); err != nil {
		errorChannel <- err
		return
	}
}

func parseByUrl(u string) (*gofeed.Feed, error) {
	fp := gofeed.NewParser()

	feed, err := fp.ParseURL(u)
	if err != nil {
		return nil, err
	}

	return feed, nil
}
