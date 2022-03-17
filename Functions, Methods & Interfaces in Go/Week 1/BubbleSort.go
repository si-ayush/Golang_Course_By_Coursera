package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Enter input numbers separated by space:=> ")
	scanner := bufio.NewReader(os.Stdin)
	input, _, _ := scanner.ReadLine()
	str := strings.Split(string(input), " ")
	var values []int
	for _, s := range str {
		n, _ := strconv.Atoi(s)
		values = append(values, n)
	}
	BubbleSort(values)
	fmt.Println(values)
}

func BubbleSort(num []int) {
	for ind := 0; ind < len(num); ind++ {
		for j := 0; j < len(num)-1-ind; j++ {
			if num[j] > num[j+1] {
				Swap(num, j)
			}
		}
	}
}

func Swap(sl []int, ind int) {
	var temp int = sl[ind]
	sl[ind] = sl[ind+1]
	sl[ind+1] = temp
}
