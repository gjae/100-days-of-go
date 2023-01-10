package main

import (
	"encoding/xml"
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Request struct {
	URL string `json:"url"`
}

type Feed struct {
	Entries []Entry `xml:"entry"`
}

type Entry struct {
	Link struct {
		Href string `xml:"href,attr"`
	} `xml:"link"`

	Thumbnail struct {
		URL string `xml:"url,attr"`
	} `xml:"media:thumbnail"`

	Author Author `xml:"author"`

	Title       string    `xml:"title"`
	PublishedAt time.Time `xml:"published"`
}

type Author struct {
	Name    string `xml:"name"`
	Profile string `xml:"uri"`
}

func GetFeedEntries(url string) ([]Entry, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}
	req.Header.Add("User-Agent", "Mozilla/5.0 (X11; Ubuntu; Linux i686; rv:24.0) Gecko/20100101 Firefox/24.0")
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	byteValue, _ := io.ReadAll(resp.Body)
	var feed Feed
	xml.Unmarshal(byteValue, &feed)
	return feed.Entries, nil
}

func ParserHandler(c *gin.Context) {
	var request Request

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	entries, err := GetFeedEntries(request.URL)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, entries)
}

func main() {
	router := gin.Default()

	router.POST("/parser", ParserHandler)

	router.Run(":5000")
}
