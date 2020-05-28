package main

import (
	"fmt"
)

func main() {
	statePopulations := map[string]int{
		"California":   39000000,
		"Texas":        27000000,
		"Florida":      20000000,
		"New York":     19000000,
		"Pennsylvania": 12000000,
		"Illinois":     12000000,
		"Ohio":         11000000,
	}
	statePopulations["Georgia"] = 10000000 //adding
	_, ok := statePopulations["Oregon"]    //check if key is present
	fmt.Println("Is state?", ok)
	fmt.Println(statePopulations["Georgia"])
}
