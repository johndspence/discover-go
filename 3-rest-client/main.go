package main

import (
  "net/http"
  "fmt"
  "encoding/json"
  "strconv"
)

func main () {
  movieid := "0372785"
  url := "http://www.omdbapi.com/?i=tt" + movieid + "&plot=short&r=json"
  resp, err := http.Get(url)
  if err != nil {
    fmt.Println("Error in file retrieval %v", err)
  }

defer resp.Body.Close()

var m Movie
if err := json.NewDecoder(resp.Body).Decode(&m); err!=nil{
  fmt.Errorf("error parsing body %v", err)
  return }

fmt.Println("status code is", resp.Status)

if m.ImdbRating == "N/A"{
  fmt.Printf("The movie : %#s was released in %#s - the IMDB rating is %#s with %#s votes\n", m.Title, m.Year, m.ImdbRating, m.ImdbVotes)
} else {
  ratingFloat, err := strconv.ParseFloat(m.ImdbRating, 64)
  ratingInt := int(ratingFloat * 10)
  if err != nil {
    fmt.Println("error", err)
    return
}
  fmt.Printf("The movie : %#s was released in %#s - the IMDB rating is %#d%% with %#s votes\n", m.Title, m.Year, ratingInt, m.ImdbVotes)
  }
}
