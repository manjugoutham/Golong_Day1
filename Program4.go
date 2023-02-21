/*Write a program which takes input of a place name where you would like to visit most and displays that with place
with uppercase once and then all lower case once.*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Print("Enter the placename : ")
	var placename string
	fmt.Scan(&placename)
	uppercase := strings.ToUpper(placename)
	lowercase := strings.ToLower(placename)
	fmt.Printf("Uppercase: %s\nLowercase: %s\n", uppercase, lowercase)

}
