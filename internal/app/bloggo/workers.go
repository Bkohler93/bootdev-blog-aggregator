package bloggo

import (
	"context"
	"database/sql"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/bkohler93/bootdev-blog-aggregator/internal/database"
	"github.com/bkohler93/bootdev-blog-aggregator/internal/feeds"
	"github.com/bkohler93/bootdev-blog-aggregator/internal/helpers"
	"github.com/google/uuid"
)

func (cfg *apiConfig) FetchFeedsWorker(numFeeds int32, interval int) {
	ticker := time.NewTicker(time.Second * time.Duration(interval))

	for {
		dbFeeds, err := cfg.DB.GetNextFeedsToFetch(context.Background(), numFeeds)
		if err != nil {
			log.Fatal(err)
		}

		wg := sync.WaitGroup{}
		for _, feed := range dbFeeds {
			wg.Add(1)
			go func(f database.Feed) {
				url := strings.Trim(f.Url, "`")
				xmlResponse := feeds.FetchFeedData(url)

				for _, item := range xmlResponse.Channel.Item {
					date := helpers.ParseDateTime(item.PubDate)

					//create post, ignore error for duplicate key error
					cfg.DB.CreatePost(context.Background(), database.CreatePostParams{
						ID:          uuid.New(),
						CreatedAt:   time.Now(),
						UpdatedAt:   time.Now(),
						Title:       item.Title,
						Url:         item.Link,
						Description: sql.NullString{String: item.Description, Valid: true},
						PublishedAt: date,
						FeedID:      f.ID,
					})
				}
			}(feed)
		}
		wg.Wait()
		<-ticker.C
	}
}
