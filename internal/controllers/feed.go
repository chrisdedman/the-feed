package controllers

import (
	"fmt"
	"html"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/microcosm-cc/bluemonday"
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

func (s *Server) GetFeed(c *gin.Context) ([]models.FeedInfo, error) {
	userID := c.MustGet("user").(jwt.MapClaims)

	var feeds []models.Feed
	var feedInfos []models.FeedInfo

	if err := s.db.Where("user_id = ?", userID["id"].(float64)).Find(&feeds).Error; err != nil {
		return nil, err
	}

	// Iterate through feeds to construct FeedInfo slice
	for _, feed := range feeds {
		feedInfos = append(feedInfos, models.FeedInfo{
			FeedID: feed.ID,
			URL:    feed.URL,
		})
	}

	return feedInfos, nil
}

func FetchFeedItems(s *Server, feedItems []models.FeedInfo) ([]FeedItem, error) {
	var allFeedItems []FeedItem
	// Create a Bluemonday policy to strip all HTML tags.
	p := bluemonday.StrictPolicy()

	for _, feedItem := range feedItems {
		feed, err := FetchFeed(feedItem.URL)
		if err != nil {
			fmt.Println("Error fetching feed:", err)
			continue
		}

		itemCount := 0
		for _, item := range feed.Items {
			if itemCount >= 10 {
				break
			}
			plainTextDescription := p.Sanitize(item.Description)
			plainTextTitle := p.Sanitize(item.Title)
			// Decode HTML entities.
			decodedDescription := html.UnescapeString(plainTextDescription)
			decodedTitle := html.UnescapeString(plainTextTitle)

			feedItem := FeedItem{
				FeedTitle:   feed.Title,
				Title:       decodedTitle,
				Description: decodedDescription,
				Link:        item.Link,
			}
			allFeedItems = append(allFeedItems, feedItem)
			itemCount++
		}
	}

	return allFeedItems, nil
}
