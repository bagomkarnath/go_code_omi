package main

import "fmt"

func main() {
	if true {
		fmt.Println("hello")
	}

	statePopulation := map[string]int{
		"odisha": 2488293,
		"andhra": 3648399,
		"bihar":  1377389,
	}

	if pop, ok := statePopulation["odisha"]; ok {
		fmt.Println(pop)
	}

	number := 30
	guess := 49

	if number < guess {
		fmt.Println("less")
	}
	if number > guess {
		fmt.Println("more")
	}
	if number == guess {
		fmt.Println("equal")
	}
	fmt.Println(number <= guess, number >= guess, number != guess)

	if number > 100 {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}

	if number > 100 {
		fmt.Println("true")
	} else if number < 100 {
		fmt.Println("false")
	} else {
		fmt.Println("else")
	}
}
