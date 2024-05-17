package main

import (
	"fmt"
)

func lowestCategory(position int) int {
	if position <= 1 {
		return 1
	} else if position <= 3 {
		return 3
	} else if position <= 5 {
		return 5
	} else if position <= 10 {
		return 10
	} else if position <= 25 {
		return 25
	} else if position <= 50 {
		return 50
	} else {
		return 100
	}
}

func main() {
	var position int
	fmt.Scan(&position)
	lowest := lowestCategory(position)
	fmt.Println("Top", lowest)
}
