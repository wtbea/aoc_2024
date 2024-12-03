package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func parse(filename string) ([]int, []int) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	var left, right []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line) // Split by whitespace

		if len(parts) == 2 {
			l, err1 := strconv.Atoi(parts[0])
			r, err2 := strconv.Atoi(parts[1])
			if err1 != nil || err2 != nil {
				panic(err1)
			}
			left = append(left, l)
			right = append(right, r)
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return left, right
}

func sort(list []int) {
	l := len(list)
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			if list[i] > list[j] {
				list[i], list[j] = list[j], list[i]
			}
		}
	}
}
