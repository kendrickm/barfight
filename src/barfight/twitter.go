package main

import (
	"github.com/ChimeraCoder/anaconda"
	"code.google.com/p/gcfg"
	"log"
	"net/url"
)

// type TwitterConfig struct {
//   Twitter struct {
//     APIKEY string
//     APISECRET string
// 		ACCESSKEY string
//     ACCESSSECRET string
//   }
// }

//Does the scraping
func PullTweets(twitter_handle string) (int64, string, string) {

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
	v.Set("count", "10")
	v.Set("contributor_details", "true")
	var Text, CreatedAt string
	var Id int64
	//log.Println(v)
	searchResult, _ := api.GetStatusesUserTimeline(v)
	for _, tweet := range searchResult {
		Text = tweet.Text
		//log.Println(tweet.Text)
		Id = tweet.Id
		CreatedAt = tweet.CreatedAt
	}
	return Id, Text, CreatedAt
}
