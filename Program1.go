// Write a program which displays 3 your favourite places names in one line,
// Enhance the program so that it displays all these names in new lines.

package main

import (
	"fmt"
)

func main() {
	var str1, str2, str3 string
	fmt.Println("Enter your 3 favourite places :")
	fmt.Scanln(&str1)
	fmt.Scanln(&str2)
	fmt.Scanln(&str3)
	fmt.Printf("favourite places names : %s,%s,%s\n", str1, str2, str3)
	fmt.Println()
	// Enhance the program so that it displays all these names in new lines.
	fmt.Printf("favourite places names : \n%s\n,%s\n,%s\n", str1, str2, str3)
}
