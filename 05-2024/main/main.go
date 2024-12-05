package main

import "fmt"

func main() {
	input := parse("input.txt")
	rules, updates := values(input)
	total := calculate(rules, updates, false)
	total_v2 := calculate(rules, updates, true)
	fmt.Println(total)
	fmt.Println(total_v2)
}
