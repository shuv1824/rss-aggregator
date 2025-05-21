package main

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/shuv1824/rss-aggregator/internal/database"
)

func startScraping(
	db database.Queries,
	concurrency int,
	timeBetweenRequest time.Duration,
) {
	log.Printf("Sraping on %v goroutines on %s duration", concurrency, timeBetweenRequest)
	ticker := time.NewTicker(timeBetweenRequest)
	defer ticker.Stop()

	for ; ; <-ticker.C {
		feeds, err := db.GetNextFeedsToFetch(context.Background(), int32(concurrency))
		if err != nil {
			log.Println("Failed to fetch feeds:", err)
			continue
		}

		wg := &sync.WaitGroup{}

		for _, feed := range feeds {
			wg.Add(1)

			go scrapeFeed(wg)
		}
		wg.Wait()
	}
}

func scrapeFeed(wg *sync.WaitGroup) {
	defer wg.Done()
}
