/*******************************************************************************
Application to scrape tap information from twitter and store in database
*******************************************************************************/

package main

import (
	"log"
	"gopkg.in/jmcvetta/napping.v1"
	"code.google.com/p/gcfg"
)

func main() {
	log.Println("*******BARFIGHT*******")
	Scraper(LoadConfig())
}

//Builds a map object of bars and their twitter handles
//func LoadConfig() map[string]string {
func LoadConfig() LocationResponse {
	var cfg ConfigFile
	cfgerr := gcfg.ReadFileInto(&cfg, "config.gcfg")
	if cfgerr != nil {
		log.Printf("Error reading config file:  %v", cfgerr)
	}
	url := cfg.BarfightAPI.URL + "/locations?type=twitter"
	s := napping.Session{}

	BarList := LocationResponse{}
	_, err := s.Get(url, nil, &BarList, nil)

	if err != nil {
		log.Printf("Error pulling twitter handles  %v", err)
	}

	return BarList

}
