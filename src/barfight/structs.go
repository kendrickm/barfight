package main

//Config file data
type ConfigFile struct {
  BarfightAPI struct {
    URL string
  }
  Twitter struct {
    APIKEY string
    APISECRET string
		ACCESSKEY string
    ACCESSSECRET string
  }
}


type CheckinRequest struct {
  Data string `json:"data"`
  Source string `json:"source"`
  Source_id string `json:"source_id"`
  Date string `json:"date"`
  Needs_review bool `json:"needs_review"`
}

type LocationResponse struct {
	Locations []SingleLocation `json:"search_results"`
}

type SingleLocation struct {
  Name string `json:"name"`
  LastScraped string `json:"last_scraped"`
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
	Source_Id string  `json:"source_id"`
}
