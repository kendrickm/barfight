package main

import (
  "github.com/fjl/go-couchdb"
  "log"
  )

func LoadConfig() map[string]string {
  c, err := couchdb.NewClient("http://docker:5984", nil)
  config_db := "config"
  if err != nil {
    log.Printf("Error creating CouchDB client:  %v", err)
  }
  db := c.DB(config_db)


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

    var result alldocsResult

    err = db.AllDocs(&result, nil)
    if err != nil {
      log.Printf("Error pulling data:  %v", err)
    }

    BarList := make(map[string]string)

    for _ , value := range result.Rows {
      s, _ := value["id"].(string)
      //BarList = append(BarList, s)
      var config configResult
      err = db.Get(s, &config, nil)
      if err != nil {
        log.Printf("Error creating CouchDB client:  %v", err)
      }
      BarList[config.Twitter] = config.Name
    }

    return BarList
}
