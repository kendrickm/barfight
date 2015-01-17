package main

import (
    // "net/http"
    // // "net/url"
    // // "github.com/ChimeraCoder/anaconda"
    // "github.com/fjl/go-couchdb"
    "log"
    // //"encoding/json"
)

func main() {
    log.Println("*******BARFIGHT*******")
    ConfigData := LoadConfig()
    for twitter, name := range ConfigData {
      log.Println(name + " is now fighting.")
      CreateDB(twitter)
      _, tweet, _ := PullTweets(twitter)
      //RawLogger(name, tweet)
      parse(tweet)
    }
}

func RawLogger(brewery string, tweet string) {
  log.Println("Raw: " + brewery + ": " + tweet)
}
