package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	ls := parse("input.txt")
	re := path(ls, false)
	re_2 := path(ls, true)
	fmt.Println(re, re_2)
	elapsed := time.Since(start).Seconds()
	fmt.Println("Total time: ", elapsed)
}
