package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func parse(filename string) [][]int {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	var lists [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		list := make([]int, len(parts))
		for i, s := range parts {
			v, err := strconv.Atoi(s)
			if err != nil {
				panic(err)
			}
			list[i] = v
		}
		lists = append(lists, list)
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return lists
}
