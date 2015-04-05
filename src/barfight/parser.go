package main

import (
	"log"
	"strings"
	"regexp"
)

func parse(rawText string) (bool, string, string) {
	needs_review := true
	rawText = stripHashtag(rawText)
	rawText = stripNames(rawText)
	rawText = stripSpecialChars(rawText)
	var (
		oldBeer = ""
		newBeer = ""
		checkin_data = ""
		)
	if strings.Contains(rawText, "replacing") {
		newBeer, oldBeer = splitOldNew(rawText)
		log.Println("Old beer is " + oldBeer)
		log.Println("New beer is " + newBeer)
		needs_review = false
		checkin_data = newBeer
	}
	if needs_review {
		log.Println("Tweet needs manual review")
		checkin_data = rawText
	}

	return needs_review, checkin_data, oldBeer
}

func stripHashtag(rawText string) string {
	var hashtagID = regexp.MustCompile(`#([^\\s]*)`)
	if hashtagID.MatchString(rawText) {
		rawText = hashtagID.ReplaceAllString(rawText, "")
	}
	return rawText
}

func stripNames(rawText string) string {
	var nameID = regexp.MustCompile(`@(\w+)`)
	if nameID.MatchString(rawText) {
		rawText = nameID.ReplaceAllString(rawText, "")
	}
	return rawText
}

func stripSpecialChars(rawText string) string {
	var openParenID = regexp.MustCompile(`\(`)
	var closeParenID = regexp.MustCompile(`\)`)
	if openParenID.MatchString(rawText) {
		rawText = openParenID.ReplaceAllString(rawText, "")
	}
	if closeParenID.MatchString(rawText) {
		rawText = closeParenID.ReplaceAllString(rawText, "")
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
