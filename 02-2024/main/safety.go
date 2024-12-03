package main

/*
The engineers are trying to figure out which reports are safe. The Red-Nosed reactor safety systems can only tolerate levels that are either gradually increasing or gradually decreasing. So, a report only counts as safe if both of the following are true:

The levels are either all increasing or all decreasing.
Any two adjacent levels differ by at least one and at most three.
In the example above, the reports can be found safe or unsafe by checking those rules:

7 6 4 2 1: Safe because the levels are all decreasing by 1 or 2.
1 2 7 8 9: Unsafe because 2 7 is an increase of 5.
9 7 6 2 1: Unsafe because 6 2 is a decrease of 4.
1 3 2 4 5: Unsafe because 1 3 is increasing but 3 2 is decreasing.
8 6 4 4 1: Unsafe because 4 4 is neither an increase or a decrease.
1 3 6 7 9: Safe because the levels are all increasing by 1, 2, or 3.
*/

func safety(list [][]int, d bool) int {
	count := 0
	for _, l := range list {
		if safe(l) {
			count++
		} else if d && dampener(l) {
			count++
		}
	}
	return count
}

func safe(li []int) bool {
	if len(li) < 2 {
		return true
	}
	increasing := li[1] > li[0]
	decreasing := li[1] < li[0]
	for i := 1; i < len(li); i++ {
		diff := li[i] - li[i-1]
		if diff < -3 || diff > 3 || diff == 0 {
			return false
		}
		if increasing && li[i] < li[i-1] {
			return false
		}
		if decreasing && li[i] > li[i-1] {
			return false
		}
	}
	return true
}

func dampener(li []int) bool {
	for i := 0; i < len(li); i++ {
		sl := make([]int, 0, len(li)-1)
		sl = append(sl, li[:i]...)
		sl = append(sl, li[i+1:]...)
		if safe(sl) {
			return true
		}
	}
	return false
}
