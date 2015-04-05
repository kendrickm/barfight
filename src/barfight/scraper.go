package main
import (
	"strconv"
	"log"
	"gopkg.in/jmcvetta/napping.v1"
	"net/http"
	"code.google.com/p/gcfg"
)

func Scraper(BarList LocationResponse){
	for _, twitter_handle := range BarList.Locations {
		log.Println(twitter_handle + " is now fighting.")
		ID, tweet, CreatedAt := PullTweets(twitter_handle)
		data := "ID: " + strconv.FormatInt(ID, 10) + " Text: " + tweet + " CreatedAt: " + CreatedAt
		RawLogger(twitter_handle, data)
		needs_review, parsedData, oldBeer := parse(tweet)

		Checkin(twitter_handle, parsedData, ID, CreatedAt, needs_review)
		if oldBeer != ""{
			Kick_Old(twitter_handle, oldBeer)
		}
	}
}



func Checkin(TwitterHandle string, BeerData string, TweetId int64, CreatedDate string, NeedsReview bool) {

	var cfg ConfigFile
	cfgerr := gcfg.ReadFileInto(&cfg, "config.gcfg")
	if cfgerr != nil {
		log.Printf("Error reading config file:  %v", cfgerr)
	}

	checkin := CheckinRequest{Data: BeerData, Source: TwitterHandle, Date: CreatedDate, Source_id: TweetId, Needs_review: NeedsReview}
	h := &http.Header{}
	h.Set("Content-Type", "application/json")
	s := napping.Session{}
	s.Header = h
	url := cfg.BarfightAPI.URL + "/checkin/twitter"

	_, err := s.Post(url, &checkin, nil, nil)
	if err != nil {
		log.Println(err)
	}
}

func Kick_Old(BarName string, OldBeer string){
	log.Println("Searching the db for " + OldBeer + " to mark kicked at bar " + BarName)
}
