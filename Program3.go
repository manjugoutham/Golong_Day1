/* Develop a Temperature Unit Converter Program. It can convert temperature from centigrade to Fahrenheit.
○ It asks user to enter temperature in degree c ○ It converts in Fahrenheit and displays the same*/

package main

import "fmt"

func main() {
	var Celcius float64 = 32.22
	Fahrenheit := 1.8*Celcius + 32
	fmt.Printf("%.2f degree Celcius is : %.2f degree Fahrenheit", Celcius, Fahrenheit)
}
