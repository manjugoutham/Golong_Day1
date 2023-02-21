/*
Compute Average of numbers. Generate 5 random numbers between 10 to 50 using "math/rand" package and compute
the average of these numbers . Print the numbers generated in one line and in next line print the average.
*/
package main

import (
	"fmt"
	"math/rand"
)

func main() {

	var total int
	total = 0
	var average float32
	for i := 0; i < 5; i++ {
		random := rand.Intn(50-10) + 10
		fmt.Printf("%v ", random)
		total = total + random
	}
	average = float32(total) / 5
	fmt.Println()
	fmt.Println("The average value is ", average)
}
