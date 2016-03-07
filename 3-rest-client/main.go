package main

import (
  "net/http"
  "fmt"
  "encoding/json"
)

func main () {

resp, err := http.Get("http://www.omdbapi.com/?i=tt0372784&plot=short&r=json")
if err != nil {
  fmt.Println("Error in file retrieval %v", err)
	// handle error
}
defer resp.Body.Close()
var m Movie

mydecoder :=json.NewDecoder(resp.Body)

if err := mydecoder.Decode(&m); err!=nil{
fmt.Errorf("error parsing body %v", err)
return }
fmt.Printf("%#v",m.Actors)


fmt.Println("status code is", resp.Status)
}
