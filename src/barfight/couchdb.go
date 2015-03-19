package main

import (
	"github.com/fjl/go-couchdb"
	"log"
	//"strings"
	// "strconv"
)
func CouchConnect() *couchdb.Client {
	c, err := couchdb.NewClient("http://docker:5984", nil)
	if err != nil {
		log.Printf("Error creating CouchDB client:  %v", err)
	}
	return c
}

func CreateDB(BarName string) {
	c := CouchConnect()
	_, err := c.EnsureDB(BarName)
	if err != nil {
		log.Printf("Error creating db:  %v", err)
	}
}

func UpdateDB(doc *checkinDocument, db string, c_id string) {
	c := CouchConnect()
	rev, err := c.DB(db).Put(c_id, doc, "")
	if err != nil {
		log.Printf("Error adding document:  %v", err)
	}
	log.Printf("Created record " + rev)
}
