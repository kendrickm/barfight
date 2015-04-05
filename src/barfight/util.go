package main
import (
	"log"
)

func RawLogger(brewery string, tweet string) {
	log.Println("Raw: " + brewery + ": " + tweet)
}
