package main
import (
	"github.com/fjl/go-couchdb"
	"log"
	"code.google.com/p/gcfg"
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
func LoadConfig() map[string]string {

	config_db := "locations" //Name of db config values are stored in
	c := CouchConnect()

	db := c.DB(config_db)
	var result alldocsResult
	BarList := make(map[string]string)

	err := db.AllDocs(&result, nil)
	if err != nil {
		log.Printf("Error pulling data:  %v", err)
	}

	//Iterates through results and builds map
	for _, value := range result.Rows {
		s, _ := value["id"].(string)
		var config configResult
		err = db.Get(s, &config, nil)
		if err != nil {
			log.Printf("Error pulling data:  %v", err)
		}

		unless config.Twitter = ""{
			BarList[config.Twitter] = config.Name
		}

	}

	return BarList
}
