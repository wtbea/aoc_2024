package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	ls := parse("input.txt")

	res := search(ls, false)
	res_v2 := search(ls, true)
	fmt.Println(res_v2)
	fmt.Println(res)
	elapsed := time.Since(start).Seconds()
	fmt.Println("Total time: ", elapsed)

}
