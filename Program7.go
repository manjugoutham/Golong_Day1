/*From a user, read his/her friend’s first name and last name. Print the full name of this friend with space
between first name and last name. */

package main

import "fmt"

func main() {

	var firstname string
	var lastname string
	fmt.Print("Enter First name for user : ")
	fmt.Scan(&firstname)
	fmt.Print("Enter Last name for user: ")
	fmt.Scan(&lastname)
	fmt.Printf("Print the fullname : %s %s", firstname, lastname)
}
