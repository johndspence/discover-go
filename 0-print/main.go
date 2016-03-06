package main

import "fmt"
import "time"

func main() {

	fmt.Println("Hello, we are Holberton School")
	p := fmt.Println
	t := time.Now()
	p(t.Format("the date is 2006-01-02 15:04:05.999999999 -07:00 MST"))
	p(t.Format("the year is 2006"))
	p(t.Format("the month is January"))
	fmt.Printf("the day is %02d\n",
		t.Day(),
	)
}
