package main
import (
	"strconv"
	"log"
	"gopkg.in/jmcvetta/napping.v1"
	"net/http"
	// "encoding/json"
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

	checkin := CheckinRequest{Data: BeerData, Source: TwitterHandle, Date: CreatedDate, Source_id: TweetId, Needs_review: NeedsReview}

	h := &http.Header{}
	h.Set("Content-Type", "application/json")
	url := "http://localhost:5000/checkin/twitter"

	//var jsonStr = []byte(`{"title":"Buy cheese and bread for breakfast."}`)


	s := napping.Session{}
	s.Header = h

	_, err := s.Post(url, &checkin, nil, nil)

	// log.Println(resp)
	if err != nil {
		log.Println(err)
	}
}

func Kick_Old(BarName string, OldBeer string){
	log.Println("Searching the db for " + OldBeer + " to mark kicked at bar " + BarName)
}
