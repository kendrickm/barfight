package main

import (
  "net/url"
  "github.com/ChimeraCoder/anaconda"
  "log"
  )

  func TwitterParse(twitter_handle string) {

    anaconda.SetConsumerKey("AnSsknpdlIDlcnA1SW9IXviGR")
    anaconda.SetConsumerSecret("GqjelkiaFVCIHeYRCTodDGbDZLcLz8q3f0gJPdatS4B03a2pog")
    api := anaconda.NewTwitterApi("16941789-VNMONPoOR2Li3qiTWVorWpm59C4TnWESIyief5IEJ",
     "dYjD2iZnGsA2pOqIG9SMUh1sus7gUkAiMLIOtHzpoBIXo")
     log.Println("Pulling tweet for " + twitter_handle)
     v := url.Values{}
     v.Set("screen_name", twitter_handle)
     v.Set("count", "1")
     v.Set("contributor_details", "true")

     //log.Println(v)
     searchResult, _ := api.GetStatusesUserTimeline(v)
     for _ , tweet := range searchResult {
       log.Println(tweet.Id)
       log.Println(tweet.CreatedAt)
       log.Println(tweet.Text)

     }
  }
