/* Get the current date n time by using time package. Print the current date time as rcvd from package
Print year, month and date separately. Also in format DD-MMM-YYYY format.*/

package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()
	fmt.Println("The current time is : ", currentTime)
	fmt.Println()
	day := time.Now().Day()
	fmt.Println("Day is      : ", day)
	weekday := time.Now().Weekday()
	fmt.Println("Day name is : ", weekday)
	month := time.Now().Month()
	fmt.Println("Month is    : ", month)
	year := time.Now().Local().Year()
	fmt.Println("Year is     : ", year)

}
