package main
import (
	"github.com/fjl/go-couchdb"
	"log"
	"code.google.com/p/gcfg"
	"gopkg.in/jmcvetta/napping.v1"
)

//Builds a couchDB object
func CouchConnect() *couchdb.Client {
	var cfg ConfigFile
	err := gcfg.ReadFileInto(&cfg, "config.gcfg")
	if err != nil {
		log.Printf("Error reading config file:  %v", err)
	}
	c, err := couchdb.NewClient(cfg.CouchDB.URL, nil)
	if err != nil {
		log.Printf("Error creating CouchDB client:  %v", err)
	}
	return c
}

//Inserts Document of Beer update
func UpdateDB(doc *checkinDocument, db string, c_id string) {
	c := CouchConnect()
	rev, err := c.DB(db).Put(c_id, doc, "")
	if err != nil {
		log.Printf("Error adding document:  %v", err)
	}
	log.Printf("Created record " + rev)
}

//Builds a map object of bars and their twitter handles
//func LoadConfig() map[string]string {
func LoadConfig() LocationResponse {
	url := "http://localhost:5000/locations?type=twitter"
	s := napping.Session{}

	BarList := LocationResponse{}
	_, err := s.Get(url, nil, &BarList, nil)

	if err != nil {
		log.Printf("Error pulling twitter handles  %v", err)
	}

	return BarList

}
