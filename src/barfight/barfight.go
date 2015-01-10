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

    DBs := CreateDBs(LoadConfig()) //Loads config data and creates db's if they're not there
    for _, name := range DBs {
      //log.Println(name)
      TwitterParse(name)

    }
    //LoadConfig()
    //log.Println(DBNames)
}
