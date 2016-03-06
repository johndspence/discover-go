package main

import (
	"fmt"
	"math/rand"
)

func main() {
	randomNumber := rand.Intn(100)
	if randomNumber > 50 {
		fmt.Printf("my random number is %d and is greater than 50\n", randomNumber)
	} else {
		fmt.Printf("my random number is %d and is less than 50\n", randomNumber)
	}
	school := "Holberton School"
	if school == "Holberton School" {
		fmt.Println("I am a student of the", school)
	}

	beautifulWeather := true
	if beautifulWeather {
		fmt.Println("It's a beautiful weather!")
	}

	holbertonFounders := []string{"Rudy Rigot", "Sylvain Kalache", "Julien Barbier"}
	for _, a := range holbertonFounders {
		fmt.Printf("%s is a founder\n", a)
	}
}
