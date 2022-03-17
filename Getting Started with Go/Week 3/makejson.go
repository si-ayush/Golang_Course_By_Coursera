/*
Write a program which prompts the user to first enter a name,
and then enter an address. Your program should create a map and
add the name and address to the map using the keys “name” and 
“address”, respectively.Your program should use Marshal() to 
create a JSON object from the map, and then your program should
print the JSON object.
*/

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {

	person := make(map[string]string)

	scanner := bufio.NewScanner(os.Stdin)
	var na string
	var add string

	fmt.Println("Enter Name: ")
	scanner.Scan()
	na = scanner.Text()

	fmt.Println("Enter Address: ")
	scanner.Scan()
	add = scanner.Text()

	person["name"] = na
	person["address"] = add

	data, e := json.MarshalIndent(person, "", " ")
	if e != nil {
		fmt.Println("Error occured! Exiting the program")
		os.Exit(0)
	} else {
		fmt.Printf("Data stored in json is %v", string(data))
	}
}
