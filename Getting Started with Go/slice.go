
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
