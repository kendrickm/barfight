package main

import (
  "github.com/fjl/go-couchdb"
  "log"
  )

//Loads the config database from couch
//Builds a map object of bars and their twitter handles

func LoadConfig() map[string]string {

  config_db := "config" //Name of db config values are stored in

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
  for _ , value := range result.Rows {
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
