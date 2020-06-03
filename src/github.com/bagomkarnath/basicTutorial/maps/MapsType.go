package main

import (
	"fmt"
)

func main() {
	statePopulation := map[string]int{
		"odisha": 2488293,
		"andhra": 3648399,
		"bihar":  1377389,
	}
	fmt.Println(statePopulation)

	statePopulation1 := make(map[string]int)
	statePopulation1 = map[string]int{
		"odisha": 2488293,
		"andhra": 3648399,
		"bihar":  1377389,
	}
	fmt.Println(statePopulation1)
	fmt.Println(statePopulation1["andhra"])
	statePopulation1["delhi"] = 778299
	fmt.Println(statePopulation1)
	delete(statePopulation1, "delhi")
	fmt.Println(statePopulation1)
	fmt.Println(statePopulation1["delhi"])
	pop, ok := statePopulation1["delhi"]
	fmt.Println(pop, ok)
	pop1, ok1 := statePopulation1["odisha"]
	fmt.Println(pop1, ok1)
}
