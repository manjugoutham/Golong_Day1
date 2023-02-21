/*
Develop a program that reads a number from user as string and print 10 times of that number as string only.
Suppose the number entered is 12 Output will be 120
*/
package main

import "fmt"

func main() {

	fmt.Println("Enter Number : ")
	var stringvalue string
	fmt.Scan(&stringvalue)
	output := stringvalue + "0"
	fmt.Println("The output value is : ", output)
}
