/*******************************************************************************
Application to scrape tap information from twitter and store in database
*******************************************************************************/

package main

import (
	"log"
)

func main() {
	log.Println("*******BARFIGHT*******")
	LocationData := LoadConfig()
	for twitter, name := range LocationData {
		log.Println(name + " is now fighting.")
		_, tweet, _ := PullTweets(twitter)
		RawLogger(name, tweet)
		parse(tweet)
	}
}

func RawLogger(brewery string, tweet string) {
	log.Println("Raw: " + brewery + ": " + tweet)
}
