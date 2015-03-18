package main

import (
	"log"
	"strings"
	"regexp"
)

func parse(rawText string) {
	needs_review := true
	rawText = stripHashtag(rawText)

	if strings.Contains(rawText, "replacing") {
		newBeer, oldBeer := splitOldNew(rawText)
		log.Println("Old beer is " + oldBeer)
		log.Println("New beer is " + newBeer)
		needs_review = false
	}
	if needs_review {
		log.Println("Tweet needs manual review")
	}
}

func stripHashtag(rawText string) string {
	var hashtagID = regexp.MustCompile(`#([^\\s]*)`)
	if hashtagID.MatchString(rawText) {
		rawText = hashtagID.ReplaceAllString(rawText, "")
	}
	return rawText
}

func splitOldNew(rawText string) (string, string) {
	splitText := strings.SplitAfter(rawText, "replacing")
	newBeer := splitText[0]
	oldBeer := splitText[1]

	newBeer = strings.Replace(newBeer, "replacing", "", -1)
	newBeer = strings.Replace(newBeer, " is on,", "", -1)
	return newBeer, oldBeer
}
