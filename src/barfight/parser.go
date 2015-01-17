package main
import (
  "log"
  "strings"
  )
func parse(rawText string) {
  needs_review := true
  if(strings.Contains(rawText, "replacing")) {
    newBeer, oldBeer := splitOldNew(rawText)
    log.Println("Old beer is " + oldBeer + " and new beer is " + newBeer)
    needs_review = false
  }
  if needs_review{
    log.Println("Tweet needs manual review")
  }
}


func splitOldNew(rawText string) (string, string) {
  splitText := strings.SplitAfter(rawText, "replacing")
  newBeer := splitText[0]
  oldBeer := splitText[1]

  return newBeer, oldBeer
}
