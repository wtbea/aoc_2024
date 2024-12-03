package main

import (
	"fmt"
	"regexp"
)

/*
It seems like the goal of the program is just to multiply some numbers. It does that with instructions like mul(X,Y), where X and Y are each 1-3 digit numbers. For instance, mul(44,46) multiplies 44 by 46 to get a result of 2024. Similarly, mul(123,4) would multiply 123 by 4.

However, because the program's memory has been corrupted, there are also many invalid characters that should be ignored, even if they look like part of a mul instruction. Sequences like mul(4*, mul(6,9!, ?(12,34), or mul ( 2 , 4 ) do nothing.

For example, consider the following section of corrupted memory:

'xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))'
Only the four highlighted sections are real mul instructions. Adding up the result of each instruction produces 161 (2*4 + 5*5 + 11*8 + 8*5).
*/

func list(lines []string) []string {
	var l []string
	p := `mul\(\d+,\d+\)`
	re := regexp.MustCompile(p)

	for _, line := range lines {
		mx := re.FindAllString(line, -1)
		l = append(l, mx...)
	}
	return l
}

func list_v2(lines []string) []string {
	var l []string
	p := `mul\(\d+,\d+\)|do\(\)|don't\(\)`
	re := regexp.MustCompile(p)

	for _, line := range lines {
		mx := re.FindAllString(line, -1)
		l = append(l, mx...)
	}
	return l
}

func add(lines []string, v2 bool) int {
	var ls []string
	if v2 {
		enabled := true
		ls = list_v2(lines)
		r := 0
		for _, s := range ls {
			if s == "do()" {
				enabled = true
				continue
			}
			if s == "don't()" {
				enabled = false
				continue
			}
			if enabled {
				r += mul(s)
			}
		}
		return r
	}
	ls = list(lines)
	r := 0
	for _, s := range ls {
		if !v2 {
			r += mul(s)
		}
	}
	return r
}

func mul(s string) int {
	var x, y int

	fmt.Sscanf(s, "mul(%d,%d)", &x, &y)

	return x * y
}
