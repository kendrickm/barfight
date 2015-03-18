package main

import (
	"github.com/ChimeraCoder/anaconda"
	"code.google.com/p/gcfg"
	"log"
	"net/url"
	"github.com/fjl/go-couchdb"
)

type Config struct {
  Twitter struct {
    APIKEY string
    APISECRET string
		ACCESSKEY string
    ACCESSSECRET string
  }
}

//Does the scraping
func PullTweets(twitter_handle string) (int64, string, string) {

	var cfg Config
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

//Builds a map object of bars and their twitter handles
func LoadConfig() map[string]string {

	config_db := "locations" //Name of db config values are stored in

	//Structs to store data returned from couch
	type alldocsResult struct {
		TotalRows int `json:"total_rows"`
		Offset    int
		Rows      []map[string]interface{}
	}
	type configResult struct {
		Rev     string `json:"_rev,omitempty"`
		Name    string `json:"name"`
		Twitter string `json:"twitter"`
	}

	c, err := couchdb.NewClient("http://docker:5984", nil)
	if err != nil {
		log.Printf("Error creating CouchDB client:  %v", err)
	}

	db := c.DB(config_db)
	var result alldocsResult
	BarList := make(map[string]string)

	err = db.AllDocs(&result, nil)
	if err != nil {
		log.Printf("Error pulling data:  %v", err)
	}

	//Iterates through results and builds map
	for _, value := range result.Rows {
		s, _ := value["id"].(string)
		var config configResult
		err = db.Get(s, &config, nil)
		if err != nil {
			log.Printf("Error creating CouchDB client:  %v", err)
		}
		BarList[config.Twitter] = config.Name
	}

	return BarList
}
