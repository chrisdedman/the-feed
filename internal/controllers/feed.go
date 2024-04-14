package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/mmcdole/gofeed"
	"github.com/the-feed/internal/models"
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

func FetchFeed(url string) (*gofeed.Feed, error) {
	feedParser := gofeed.NewParser()
	feed, err := feedParser.ParseURL(url)
	if err != nil {
		return nil, fmt.Errorf("error parsing feed from %s: %v", url, err)
	}

	return feed, nil
}

func (s *Server) AddFeed(c *gin.Context) {
	userID := c.MustGet("user").(jwt.MapClaims)
	fmt.Println("User ID:", userID)

	var user models.User
	if err := s.db.Where("id = ?", userID["id"].(float64)).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not authorized to add feed"})
		return
	}

	var input UserFeed
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	url := models.Feed{URL: input.Url, UserID: userID["id"].(float64)}

	if err := s.db.Create(&url).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "feed added successfully"})
}

func (s *Server) GetFeed(c *gin.Context) ([]models.Feed, error) {
	userID := c.MustGet("user").(jwt.MapClaims)

	var feed []models.Feed

	// Retrieve the feed associated with the user ID
	if err := s.db.Where("user_id = ?", userID["id"].(float64)).Find(&feed).Error; err != nil {
		return nil, err
	}

	return feed, nil
}

func FetchFeedItems(feedItems []models.Feed) ([]FeedItem, error) {
	var allFeedItems []FeedItem
	for _, feedItem := range feedItems {
		feed, err := FetchFeed(feedItem.URL)
		if err != nil {
			fmt.Println("Error fetching feed:", err)
			continue // Skip to the next feed
		}

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
