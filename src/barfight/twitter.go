package main

import (
	"github.com/ChimeraCoder/anaconda"
	"code.google.com/p/gcfg"
	"log"
	"net/url"
	"strconv"
)

//Does the scraping
func PullTweets(twitter_handle string, since_id string) {

	var cfg ConfigFile
	err := gcfg.ReadFileInto(&cfg, "config.gcfg")
	if err != nil {
		log.Printf("Error reading config file:  %v", err)
	}
	anaconda.SetConsumerKey(cfg.Twitter.APIKEY)
	anaconda.SetConsumerSecret(cfg.Twitter.APISECRET)
	api := anaconda.NewTwitterApi(cfg.Twitter.ACCESSKEY, cfg.Twitter.ACCESSSECRET)
	log.Println("Pulling tweet for " + twitter_handle)
	v := url.Values{}
	v.Set("screen_name", twitter_handle)
	v.Set("contributor_details", "true")
	v.Set("since_id", since_id)
	searchResult, _ := api.GetStatusesUserTimeline(v)

	if len(searchResult) == 0 {
		log.Println("No new tweets")
	}	else {
		for i := len(searchResult)-1; i >= 0; i-- {
	    tweet := searchResult[i]
			id := strconv.FormatInt(tweet.Id, 10)
			WorkScrapings(twitter_handle, id, tweet.Text, tweet.CreatedAt)
		}
		//
		// for _, tweet := range searchResult {
		// 	id := strconv.FormatInt(tweet.Id, 10)
		// 	WorkScrapings(twitter_handle, id, tweet.Text, tweet.CreatedAt)
		// }
	}
}
