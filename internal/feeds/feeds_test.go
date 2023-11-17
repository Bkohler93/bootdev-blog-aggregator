package feeds

import (
	"testing"
)

func TestFetchFeedData(t *testing.T) {
	testUrl := "https://blog.boot.dev/index.xml"

	data := FetchFeedData(testUrl)

	if data.Channel.Title != "Boot.dev Blog" {
		t.Errorf("Expected title to be 'Boot.dev Blog', got %s", data.Channel.Title)
	}
}

func TestFetchFeedData2(t *testing.T) {
	testUrl := "https://wagslane.dev/index.xml"

	data := FetchFeedData(testUrl)

	if data.Channel.Title != "Lane's Blog" {
		t.Errorf("Expected title to be 'Lanes Blog', got %s", data.Channel.Title)
	}
}
