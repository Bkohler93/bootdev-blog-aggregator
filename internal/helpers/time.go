package helpers

import (
	"log"
	"time"
)

func ParseDateTime(s string) time.Time {
	layouts := []string{"Mon, 02 Jan 2006 15:04:05 -0700",
		"Mon, 02 Jan 2006 15:04:05 MST"}

	var parsedTime time.Time
	var err error
	var isParsed bool

	for _, layout := range layouts {
		parsedTime, err = time.Parse(layout, s)
		if err != nil {
			continue
		} else {
			isParsed = true
			break
		}
	}

	if !isParsed {
		log.Println("could not parse time", err)
	}
	return parsedTime
}
