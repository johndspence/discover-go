package main

import "fmt"

type user struct {
	Name string
	DOB  string
	City string
}

func (u *user) sayHello() {
	fmt.Printf("Hey! %s\n", u.Name)
}

func (u *user) saySentence() {
	fmt.Printf("%s who was born in %s would be 99 years old today.\n.", u.Name, u.City)
}

func main() {
	u := user{Name: "Betty Holberton", DOB: "March 7, 1917", City: "Philadelphia"}
	u.sayHello()
	u.saySentence()

//   func Parse(Mon Jan 2 15:04:05 -0700 MST 2006, value u.DOB) (Time, error)
// t equals parsed DOB

func Since(Mon Jan 2 15:04:05 -0700 MST 2006 Time) Hour

}

// time.now

  // https://golang.org/pkg/time/
  // func Since(t Time) Duration

	// d:=dog{Breed:"GermanShepherd",Age:14,Name:"Clyde"}
	// b,err:=json.Marshal(d)
	// iferr!=nil{
	// fmt.Printf("can'tmarshal%v",err) }
	// fmt.Println(string(b))
	//
	// 	u := user{Name: "Betty Holberton", DOB: "March 7, 1917", City: "Philadelphia"}
	//   b, err := json.Marshal(u)
	//
	//   .fmt.Printf("Hello %s\n", u.Name)


// 	randomNumber := rand.Intn(100)
// 	if randomNumber > 50 {
// 		fmt.Printf("my random number is %d and is greater than 50\n", randomNumber)
// 	} else {
// 		fmt.Printf("my random number is %d and is less than 50\n", randomNumber)
// 	}
// 	school := "Holberton School"
// 	if school == "Holberton School" {
// 		fmt.Println("I am a student of the", school)
// 	}
//
// 	beautifulWeather := true
// 	if beautifulWeather {
// 		fmt.Println("It's a beautiful weather!")
// 	}
//
// 	holbertonFounders := []string{"Rudy Rigot", "Sylvain Kalache", "Julien Barbier"}
// 	for _, a := range holbertonFounders {
// 		fmt.Printf("%s is a founder\n", a)
// 	}
// }
