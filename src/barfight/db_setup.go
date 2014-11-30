package main

import (
  "github.com/fjl/go-couchdb"
  "log"
  "strings"
  )

func CreateDBs(BarList map[string]string) []string {

  c, err := couchdb.NewClient("http://docker:5984", nil)
  if err != nil {
    log.Printf("Error creating CouchDB client:  %v", err)
  }

  DBNames := make([]string, 3)

  for twitter, name := range BarList {
    //log.Println("Current name is " + name + " and twitter is " + twitter)
    CleanName := strings.Replace(name, " ", "", -1)
    CleanName = strings.ToLower(CleanName)
    log.Println("Updated name is " + CleanName + " and twitter is " + twitter)
    _, err := c.EnsureDB(CleanName)
    if err != nil {
      log.Printf("Error creating db:  %v", err)
    }
  }
}
