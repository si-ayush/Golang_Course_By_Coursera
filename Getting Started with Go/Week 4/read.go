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
