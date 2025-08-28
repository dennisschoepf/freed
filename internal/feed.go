package internal

import (
	"fmt"
	"freed/internal/database"
	"net/url"
	"strconv"
	"sync"
	"time"

	"github.com/mmcdole/gofeed"
	"github.com/pterm/pterm"
)

func AddFeed(feedUrl string) (string, int, error) {
	if _, err := url.ParseRequestURI(feedUrl); err != nil {
		return "", 0, fmt.Errorf("The given URL does not seem to be valid: %s", err)
	}

	feed, err := parseByUrl(feedUrl)

	if err != nil {
		return "", 0, err
	}

	f := database.Feed{
		Name: feed.Title,
		Url:  feedUrl,
	}

	feedId, err := f.Insert()
	if err != nil {
		return "", 0, err
	}

	// TODO: Make the amount of articles configurable
	articles := make([]database.Article, 0, 10)

	for i, v := range feed.Items {
		if i > 10 {
			continue
		}

		a := database.Article{
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

func GetAllFeedsAsTable() (pterm.TableData, error) {
	feeds, err := database.FindAllFeedsWithArticleCount()

	if err != nil {
		return nil, err
	}

	tableData := pterm.TableData{}
	headerRow := []string{"ID", "Name", "Url", "Item Count", "Created At", "Last Synced At"}

	tableData = append(tableData, headerRow)

	for _, feed := range *feeds {
		rowData := []string{
			fmt.Sprintf("[%d]", feed.ID),
			feed.Name,
			feed.Url,
			strconv.Itoa(feed.ArticleCount),
			pterm.Gray(feed.CreatedAt),
			pterm.Gray(feed.LastSyncedAt),
		}

		tableData = append(tableData, rowData)
	}

	return tableData, nil
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
	feed database.Feed,
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

	var articles []database.Article

	for _, item := range gFeed.Items {
		if item.PublishedParsed.Before(*feed.LastSyncedAt) {
			continue
		}

		article := database.Article{
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
