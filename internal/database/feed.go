package database

import "fmt"

func AddFeed(url string) error {
	fmt.Printf("Trying to add new feed by URL: %s", url)
	return nil
}
