package main

import (
  "github.com/fjl/go-couchdb"
  "log"
  "strings"
  //"strconv"
  )

func CreateDBs(BarList map[string]string) []string {

  c, err := couchdb.NewClient("http://docker:5984", nil)
  if err != nil {
    log.Printf("Error creating CouchDB client:  %v", err)
  }

//  log.Println("Creating array of " + strconv.Itoa(len(BarList)) + " size.")
  DBNames := make([]string, 0)

  for twitter, name := range BarList {
    //log.Println("Current name is " + name + " and twitter is " + twitter)
    CleanName := strings.Replace(name, " ", "", -1)
    CleanName = strings.ToLower(CleanName)
    log.Println("Updated name is " + CleanName + " and twitter is " + twitter)
    DBNames = append(DBNames, CleanName)
    _, err := c.EnsureDB(CleanName)
    if err != nil {
      log.Printf("Error creating db:  %v", err)
    }
  }
  // log.Println("Final size of array is " + strconv.Itoa(len(DBNames)))
  return DBNames
}


func UpdateDB(BarName string, Info string) {

}
