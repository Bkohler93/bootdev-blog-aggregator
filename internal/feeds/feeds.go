package feeds

import (
	"encoding/xml"
	"log"
	"net/http"
)

type XMLResponse struct {
	XMLName xml.Name `xml:"rss"`
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Language    string    `xml:"language"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func FetchFeedData(url string) XMLResponse {
	data := XMLResponse{}
	client := http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	res, err := client.Do(req)
	if err != nil {
		log.Println("error executing request", err)
	}

	decoder := xml.NewDecoder(res.Body)
	decoder.Decode(&data)
	return data
}
