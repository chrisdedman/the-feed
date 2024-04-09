package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-web-app/internal/models"
	"github.com/mmcdole/gofeed"
)

// UserFeed represents a user input for a feed.
type UserFeed struct {
	Url string `json:"url" binding:"required"`
}

// FeedItem represents a single feed item
type FeedItem struct {
	FeedTitle   string `json:"feedTitle"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Link        string `json:"link"`
}

// fetchFeed fetches a feed from the given URL and returns the parsed feed.
func FetchFeed(url string) (*gofeed.Feed, error) {
	feedParser := gofeed.NewParser()
	feed, err := feedParser.ParseURL(url)
	if err != nil {
		return nil, fmt.Errorf("error parsing feed from %s: %v", url, err)
	}

	return feed, nil
}

// AddFeed add the feed link to the database
func (s *Server) AddFeed(c *gin.Context) {
	var input UserFeed
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println(input.Url)

	url := models.Feed{URL: input.Url}

	if err := s.db.Create(&url).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "feed added successfully"})
}

// GetFeed gets the feed from the database
func (s *Server) GetFeed(c *gin.Context) ([]models.Feed, error) {
	var feed []models.Feed

	if err := s.db.Find(&feed).Error; err != nil {
		return nil, err
	}

	return feed, nil
}

// FetchFeedItems fetches feed items from the given URL and returns them.
func FetchFeedItems(feedItems []models.Feed) ([]FeedItem, error) {
	var allFeedItems []FeedItem
	for _, feedItem := range feedItems {
		feed, err := FetchFeed(feedItem.URL)
		if err != nil {
			fmt.Println("Error fetching feed:", err)
			continue // Skip to the next feed
		}

		// Iterate over feed items and convert them to controllers.FeedItem
		for _, item := range feed.Items {
			feedItem := FeedItem{
				FeedTitle:   feed.Title,
				Title:       item.Title,
				Description: item.Description,
				Link:        item.Link,
			}
			allFeedItems = append(allFeedItems, feedItem)
		}
	}
	return allFeedItems, nil
}
