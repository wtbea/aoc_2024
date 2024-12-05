package main

import (
	"sort"
	"strconv"
	"strings"
)

/*
In the above example, the first update (75,47,61,53,29) is in the right order:

75 is correctly first because there are rules that put each other page after it: 75|47, 75|61, 75|53, and 75|29.
47 is correctly second because 75 must be before it (75|47) and every other page must be after it according to 47|61, 47|53, and 47|29.
61 is correctly in the middle because 75 and 47 are before it (75|61 and 47|61) and 53 and 29 are after it (61|53 and 61|29).
53 is correctly fourth because it is before page number 29 (53|29).
29 is the only page left and so is correctly last.
Because the first update does not include some page numbers, the ordering rules involving those missing page numbers are ignored.

The second and third updates are also in the correct order according to the rules. Like the first update, they also do not include every page number, and so only some of the ordering rules apply - within each update, the ordering rules that involve missing page numbers are not used.

The fourth update, 75,97,47,61,53, is not in the correct order: it would print 75 before 97, which violates the rule 97|75.

The fifth update, 61,13,29, is also not in the correct order, since it breaks the rule 29|13.

The last update, 97,13,75,29,47, is not in the correct order due to breaking several rules.

For some reason, the Elves also need to know the middle page number of each update being printed. Because you are currently only printing the correctly-ordered updates, you will need to find the middle page number of each correctly-ordered update. In the above example, the correctly-ordered updates are:

75,47,61,53,29
97,61,53,29,13
75,29,13
*/

func values(lines []string) (map[int][]int, [][]int) {
	rules := map[int][]int{}
	updates := [][]int{}
	parse_updates := false
	for _, line := range lines {
		if line == "" {
			parse_updates = true
			continue
		}
		if parse_updates {
			updates = append(updates, update(line))
		} else {
			before, after := rule(line)
			rules[before] = append(rules[before], after)
		}
	}
	return rules, updates
}

func rule(s string) (int, int) {
	parts := strings.Split(s, "|")
	before, _ := strconv.Atoi(parts[0])
	after, _ := strconv.Atoi(parts[1])
	return before, after
}

func update(s string) []int {
	res := []int{}
	parts := strings.Split(s, ",")
	for _, part := range parts {
		val, _ := strconv.Atoi(part)
		res = append(res, val)
	}
	return res
}

func valid(r map[int][]int, update []int) bool {
	for i := 0; i < len(update)-1; i++ {
		next := update[i+1]
		current := update[i]
		if contains(r[next], current) {
			return false
		}
	}
	return true
}

func contains(rules []int, val int) bool {
	for _, rule := range rules {
		if rule == val {
			return true
		}
	}
	return false
}

func calculate(r map[int][]int, updates [][]int, v2 bool) int {
	sum := 0
	for _, update := range updates {
		if v2 {
			/*
				For each of the incorrectly-ordered updates,
				use the page ordering rules to put the page numbers in the right order.
				75,97,47,61,53 becomes 97,75,47,61,53.
				61,13,29 becomes 61,29,13.
				97,13,75,29,47 becomes 97,75,47,29,13.
				After taking only the incorrectly-ordered updates and ordering them correctly,
				their middle page numbers are 47, 29, and 47. Adding these together produces 123.
			*/
			if valid(r, update) {
				continue
			}
			sort.Slice(update, func(i, j int) bool {
				return contains(r[update[i]], update[j])
			})
			sum += update[len(update)/2]
		} else {
			if valid(r, update) {
				sum += update[len(update)/2]
			}
		}

	}
	return sum
}
