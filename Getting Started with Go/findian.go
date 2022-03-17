package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	inp := scanner.Text()
	// fmt.Println(inp)

	s := strings.ToLower(inp)
	t := strings.ReplaceAll(s, " ", "")

	if strings.HasPrefix(t, "i") {
		if strings.HasSuffix(t, "n") {
			if strings.Contains(t, "a") {
				fmt.Println("Found!")
			} else {
				fmt.Println("Not Found!")
			}
		} else {
			fmt.Println("Not Found!")
		}
	} else {
		fmt.Println("Not Found!")
	}

}
