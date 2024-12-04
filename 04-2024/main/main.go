package main

import "fmt"

func main() {
	ls := parse("input.txt")

	res := search(ls, false)
	res_v2 := search(ls, true)

	fmt.Println(res_v2)
	fmt.Println(res)
}
