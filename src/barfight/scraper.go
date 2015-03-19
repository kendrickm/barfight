package main
import (
//	"github.com/fjl/go-couchdb"
	// "log"
	//"strings"
	"strconv"
)


//Structs used to send new document for checkin
type checkinDocument struct {
	Rev       string `json:"_rev,omitempty"`
	Beer      string `json:"beer"`
	Date      string `json:"date"`
  Source    string `json:"source"`
	Source_Id int64  `json:"source_id"`
}

func Checkin(BarName string, BeerData string, TweetId int64, CreatedAt string, NeedsReview bool) {
  var DB = ""
  if NeedsReview{
    DB = "checkins_needreview"
  } else {
    DB = "checkins"
  }
  doc := &checkinDocument{Beer: BeerData, Source: "twitter", Source_Id: TweetId, Date: CreatedAt}
	c_id := strconv.FormatInt(TweetId, 10)
  UpdateDB(doc, DB, c_id)
}
