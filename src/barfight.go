package main

import (
    "fmt"
    "net/http"
    // "net/url"
    // "github.com/ChimeraCoder/anaconda"
    "github.com/fjl/go-couchdb"
    "log"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
    config_db := "config"
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

    c, err := couchdb.NewClient("http://docker:5984", nil)
    if err != nil {
      log.Printf("Error creating CouchDB client:  %v", err)
    }

    db, _ := c.CreateDB(config_db)
    
    // if err != nil {
    //   if err. == "PUT http://docker:5984/foo: (412) file_exists: The database could not be created, the file already exists." {
    //     log.Printf("DB already exists; moving on")
    //   } else {
    //     log.Printf("Error creating couchdb db:  %v", err)
    //   }
    // }
}
