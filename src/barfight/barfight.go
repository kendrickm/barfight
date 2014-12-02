package main

import (
    // "net/http"
    // // "net/url"
    // // "github.com/ChimeraCoder/anaconda"
    // "github.com/fjl/go-couchdb"
    //"log"
    // //"encoding/json"
)

func main() {

    // anaconda.SetConsumerKey("AnSsknpdlIDlcnA1SW9IXviGR")
    // anaconda.SetConsumerSecret("GqjelkiaFVCIHeYRCTodDGbDZLcLz8q3f0gJPdatS4B03a2pog")
    // api := anaconda.NewTwitterApi("16941789-VNMONPoOR2Li3qiTWVorWpm59C4TnWESIyief5IEJ",
    //  "dYjD2iZnGsA2pOqIG9SMUh1sus7gUkAiMLIOtHzpoBIXo")
    //
    //  v := url.Values{}
    //  v.Set("screen_name", "thehopandvine")
    //  searchResult, _ := api.GetStatusesUserTimeline(v)
    //  for _ , tweet := range searchResult {
    //    fmt.Println(tweet.Text)
    //  }

    DBNames := CreateDBs(LoadConfig()) //Loads config data and creates db's if they're not there
    
}
