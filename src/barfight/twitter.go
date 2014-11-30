package main

import (
  "net/http"
  "github.com/ChimeraCoder/anaconda"
  "log"
  )

  func TwitterParse() {

    anaconda.SetConsumerKey("AnSsknpdlIDlcnA1SW9IXviGR")
    anaconda.SetConsumerSecret("GqjelkiaFVCIHeYRCTodDGbDZLcLz8q3f0gJPdatS4B03a2pog")
    api := anaconda.NewTwitterApi("16941789-VNMONPoOR2Li3qiTWVorWpm59C4TnWESIyief5IEJ",
     "dYjD2iZnGsA2pOqIG9SMUh1sus7gUkAiMLIOtHzpoBIXo")

     v := url.Values{}
     v.Set("screen_name", "thehopandvine")
     searchResult, _ := api.GetStatusesUserTimeline(v)
     for _ , tweet := range searchResult {
       fmt.Println(tweet.Text)
     }
  }
