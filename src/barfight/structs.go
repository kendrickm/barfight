package main

//Config file data
type ConfigFile struct {
  CouchDB struct {
    URL string
  }
  Twitter struct {
    APIKEY string
    APISECRET string
		ACCESSKEY string
    ACCESSSECRET string
  }
}

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

//Structs used to send new document for checkin
type checkinDocument struct {
	Rev       string `json:"_rev,omitempty"`
	Beer      string `json:"beer"`
	Date      string `json:"date"`
  Source    string `json:"source"`
	Source_Id int64  `json:"source_id"`
}
