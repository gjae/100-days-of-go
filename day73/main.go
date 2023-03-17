package main

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type Rss struct {
	XMLName xml.Name  `xml:"rss"`
	Text    string    `xml:",chardata"`
	Version string    `xml:"version,attr"`
	Sparkle string    `xml:"sparkle,attr"`
	Channel []Channel `xml:"channel"`
}

type Channel struct {
	Text  string `xml:",chardata"`
	Title string `xml:"title"`
	Item  Item   `xml:"item"`
}

type Item struct {
	Text        string `xml:",chardata"`
	Title       string `xml:"title"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
	Enclosure   struct {
		Text    string `xml:",chardata"`
		URL     string `xml:"url,attr"`
		Version string `xml:"version,attr"`
		Os      string `xml:"os,attr"`
	} `xml:"enclosure"`
}

func main() {
	if len(os.Args) < 2 {
		fmt.Errorf("Error")
		return
	}

	request, err := http.Get(os.Args[1])
	var ver Rss
	if err != nil {
		panic(err)
	}

	defer request.Body.Close()

	if err := xml.NewDecoder(request.Body).Decode(&ver); err != nil {
		panic(err)
	}

	verParts := strings.Split(ver.Channel[0].Item.Enclosure.Version, ".")
	mayor, _ := strconv.ParseFloat(strings.Join(verParts[:2], "."), 32)
	minor, _ := strconv.ParseInt(verParts[2], 10, 64)

	fmt.Println(mayor, minor)
	fmt.Println(ver.Channel[0].Title, "\n", ver.Channel[0].Item.Description, "\n", ver.Channel[0].Item.Enclosure.Os, "\n", ver.Channel[0].Item.Enclosure.Version)

}
