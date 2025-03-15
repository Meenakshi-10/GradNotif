package scrape

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/vartanbeno/go-reddit/v2/reddit"
)

// Time limit for filtering recent posts (60 minutes)
const timeLimit = 60 * time.Minute

// Check if a post mentions target universities and return the university name if found
func mentionsTargetUniversity(title, body string) (bool, string) {
	universities := strings.Split(os.Getenv("TARGET_UNIVERSITIES"), ",")
	for _, uni := range universities {
		fmt.Println(uni)
		if contains(title, uni) || contains(body, uni) {
			return true, uni
		}
	}
	return false, ""
}

// Helper function to check if a string contains a substring (case-insensitive)
func contains(text, substring string) bool {
	return strings.Contains(strings.ToLower(text), strings.ToLower(substring))
}

// ScrapeSubreddit fetches and sends relevant posts from a given subreddit.
func ScrapeSubreddit(client *reddit.Client, subreddit string, limit int, mentionsCh chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()

	// Fetch recent posts from the subreddit
	posts, _, err := client.Subreddit.NewPosts(context.Background(), subreddit, &reddit.ListOptions{
		Limit: limit,
	})

	if err != nil {
		log.Printf("Error fetching posts from %s: %v", subreddit, err)
		return
	}

	now := time.Now()

	for _, post := range posts {
		postTime := post.Created.Time // Access the creation time correctly
		if now.Sub(postTime) <= timeLimit {
			found, university := mentionsTargetUniversity(post.Title, post.Body)
			if found {
				mention := fmt.Sprintf("Mention: %s. Link: %s", university, post.URL)
				mentionsCh <- mention
			}
		}
	}
}

// Scrape collects mentions from multiple subreddits concurrently and returns them as an array.
func Scrape() []string {
	client, err := reddit.NewReadonlyClient()
	if err != nil {
		log.Fatalf("Error creating Reddit client: %v", err)
	}

	// Subreddits to scrape
	subreddits := []string{"gradadmissions", "mscs"}
	limit := 20

	// Use a channel to collect mentions
	mentionsCh := make(chan string)
	var wg sync.WaitGroup

	for _, subreddit := range subreddits {
		wg.Add(1)
		go ScrapeSubreddit(client, subreddit, limit, mentionsCh, &wg)
	}

	// Close the channel once all scraping is done
	go func() {
		wg.Wait()
		close(mentionsCh)
	}()

	// Collect all mentions into an array
	var mentions []string
	for mention := range mentionsCh {
		mentions = append(mentions, mention)
	}

	fmt.Println("Scraping completed.")
	return mentions
}
