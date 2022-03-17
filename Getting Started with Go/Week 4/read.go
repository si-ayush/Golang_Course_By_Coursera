/*
Write a program which reads information from a file and
represents it in a slice of structs. Assume that there 
is a text file which contains a series of names. Each 
line of the text file has a first name and a last name,
in that order, separated by a single space on the line. 

Your program will define a name struct which has two fields,
fname for the first name, and lname for the last name. Each
field will be a string of size 20 (characters).

Your program should prompt the user for the name of the text
file. Your program will successively read each line of the 
text file and create a struct which contains the first and 
last names found in the file. Each struct created will be 
added to a slice, and after all lines have been read from 
the file, your program will have a slice containing one struct 
for each line in the file. After reading all lines from the file,
your program should iterate through your slice of structs and 
print the first and last names found in each struct.
*/


package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Name struct {
	first string
	last  string
}

func main() {
	sl := make([]Name, 0, 1)
	var file string

	fmt.Println("Enter file name (example-- sample.txt. Write file name with format):")
	fmt.Scan(&file)
	fmt.Printf("File name is %v", file)
	fmt.Printf("\n")

	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), " ")
		var temp Name
		temp.first = s[0]
		temp.last = s[1]
		sl = append(sl, temp)
	}

	fmt.Println("Names read from the file are: ")
	for ind, n := range sl {
		fmt.Printf(" %v  First name is: %v,  Last name is: %v \n", ind, n.first, n.last)

	}
}
