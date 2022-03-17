package main

import "fmt"

func main() {
	var num float64
	var result int64
	fmt.Println("Enter a floating piont number")
	fmt.Scan(&num)
	result = int64(num)
	fmt.Println("Entered number is ", num)
	fmt.Println("Truncated number is ", result)
}
