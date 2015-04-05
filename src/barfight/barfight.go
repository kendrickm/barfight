/*******************************************************************************
Application to scrape tap information from twitter and store in database
*******************************************************************************/

package main

import (
	"log"
)

func main() {
	log.Println("*******BARFIGHT*******")
	//LocationData := LoadConfig()
	Scraper(LoadConfig())
	//LoadConfig()
}
