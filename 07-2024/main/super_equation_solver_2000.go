package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parse(filename string) []string {
	var lines []string

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return lines
}

/*
3267: 81 40 27
key will be 3267
value will be a listr with 81 40 27
*/
func values(ls []string) map[int][]int {
	m := map[int][]int{}

	for _, l := range ls {
		parts := strings.SplitN(l, ":", 2)
		if len(parts) != 2 {
			continue
		}
		key, err := strconv.Atoi(strings.TrimSpace(parts[0]))
		if err != nil {
			continue
		}
		valuesStr := strings.Fields(parts[1])
		values := []int{}
		for _, v := range valuesStr {
			num, err := strconv.Atoi(v)
			if err != nil {
				continue
			}
			values = append(values, num)
		}
		m[key] = values
	}
	return m
}

func evaluate(numbers []int, ops []byte, v2 bool) int {
	result := numbers[0]
	for i, op := range ops {
		if op == '+' {
			result += numbers[i+1]
		} else if op == '*' {
			result *= numbers[i+1]
		} else if v2 && op == '|' {
			result = concatenate(result, numbers[i+1])
		}
	}
	return result
}

func concatenate(x, y int) int {
	strX := strconv.Itoa(x)
	strY := strconv.Itoa(y)
	concatenated, _ := strconv.Atoi(strX + strY)
	return concatenated
}

func match(testValue int, numbers []int, ops []byte, pos int, v2 bool) bool {
	if pos == len(numbers)-1 {
		return evaluate(numbers, ops, v2) == testValue
	}
	ops[pos] = '+'
	if match(testValue, numbers, ops, pos+1, v2) {
		return true
	}
	ops[pos] = '*'
	if match(testValue, numbers, ops, pos+1, v2) {
		return true
	}
	if v2 {
		ops[pos] = '|'
		if match(testValue, numbers, ops, pos+1, v2) {
			return true
		}
	}
	return false
}

func total(eqs map[int][]int, v2 bool) int {
	total := 0
	for testValue, numbers := range eqs {
		ops := make([]byte, len(numbers)-1)
		if match(testValue, numbers, ops, 0, v2) {
			total += testValue
		}
	}
	return total
}

func test() {
	ls := parse("input.txt")
	m := values(ls)
	fmt.Println(total(m, false))
	fmt.Println(total(m, true))
}
