/*
Write a program which prompts the user to enter integers
and stores the integers in a sorted slice. The program
shouldbe written as a loop. Before entering the loop,
the program should create an empty integer slice of
size (length) 3. During each pass through the loop, the
program prompts the user to enter an integer to be
added to the slice. The program adds the integer to the
slice, sorts the slice, and prints the contents of the
slice in sorted order. The slice must grow in size to
accommodate any number of integers which the user decides
to enter. The program should only quit (exiting the loop)
when the user enters the character ‘X’ instead of an integer
*/

package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	var inp string
	var sl = make([]int, 3)
	var t bool = true

	for t {

		fmt.Println("Enter the number, press X to exit")

		fmt.Scanln(&inp)

		if inp == "X" {
			t = false
			break
		}
		n, _ := strconv.Atoi(inp)
		ind, f := findPosition(sl, 0)
		if f {
			sl[ind] = n
		} else {
			sl = append(sl, n)
		}
		sort.Ints(sl)
		//fmt.Println(sl)
		fmt.Printf("Now slice elements are: %v\n", sl)
	}
}

func findPosition(sl []int, b int) (int, bool) {
	for ind, val := range sl {
		if val == b {
			return ind, true
		}
	}

	return -1, false
}
