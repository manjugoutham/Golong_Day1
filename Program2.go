// Write a program which takes diameter (float32) of a circle and computes perimeter and area of the circle
package main

import (
	"fmt"
)

func main() {

	const pi float32 = 3.14
	var diameter float32
	fmt.Println("Enter Diameter value:")
	fmt.Scan(&diameter)
	perimeter := pi * diameter
	fmt.Printf("Perimeter for Diameter %v is %v", diameter, perimeter)
	fmt.Println()
	area := 3.14 * (diameter / 2) * (diameter / 2)
	fmt.Println("The are is : ", area)
}
