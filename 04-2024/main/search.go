package main

/*
This word search allows words to be horizontal, vertical, diagonal, written backwards, or even overlapping
other words. It's a little unusual, though, as you don't merely need to find one instance of XMAS -
you need to find all of them. Here are a few ways XMAS might appear,
where irrelevant characters have been replaced with .:
*/

func search(input []string, v2 bool) int {
	ch := 'X'
	if v2 {
		ch = 'A'
	}

	c := 0
	for x := 0; x < len(input); x++ {
		for y, char := range input[x] {
			if char == ch {
				if v2 {
					c += check_v2(input, x, y)
				} else {
					c += check(input, x, y)
				}
			}
		}
	}
	return c
}

func check(input []string, x, y int) int {
	c := 0
	length := 4

	// horizontal
	if y >= length-1 && match(input, x, y, 0, -1) {
		c++
	}
	if y <= len(input[x])-length && match(input, x, y, 0, 1) {
		c++
	}
	// vertical
	if x >= length-1 && match(input, x, y, -1, 0) {
		c++
	}
	if x <= len(input[x])-length && match(input, x, y, 1, 0) {
		c++
	}
	// diagonal
	if x >= length-1 && y <= len(input[x])-length && match(input, x, y, -1, 1) {
		c++
	}
	if x >= length-1 && y >= length-1 && match(input, x, y, -1, -1) {
		c++
	}
	if x <= len(input)-length && y <= len(input[x])-length && match(input, x, y, 1, 1) {
		c++
	}
	if x <= len(input)-length && y >= length-1 && match(input, x, y, 1, -1) {
		c++
	}

	return c
}

func match(input []string, x, y, dx, dy int) bool {
	// for this it will be (x, y) as starting point and (dx, dy) as direction
	// and it will look for XMAS as a pattern
	const p = "XMAS"
	for k := 0; k < len(p); k++ {
		if input[x+dx*k][y+dy*k] != p[k] {
			return false
		}
	}
	return true
}

/*
Looking for the instructions, you flip over the word search to find that this isn't actually an XMAS puzzle; it's an X-MAS puzzle in which you're supposed to find two MAS in the shape of an X. One way to achieve that is like this:

M.S
.A.
M.S
Irrelevant characters have again been replaced with . in the above diagram. Within the X, each MAS can be written forwards or backwards.

Here's the same example from before, but this time all of the X-MASes have been kept instead:

.M.S......
..A..MSMS.
.M.S.MAA..
..A.ASMSM.
.M.S.M....
..........
S.S.S.S.S.
.A.A.A.A..
M.M.M.M.M.
..........
*/

func check_v2(input []string, x, y int) int {
	c := 0

	if x-1 < 0 || x+1 >= len(input) || y-1 < 0 || y+1 >= len(input[x]) {
		return 0
	}

	/*
		  1:    	   2:    		3:    		4:
		M   M        S   S        M   S        S   M
		  A            A            A            A
		S   S        M   M        M   S        S   M
	*/

	configs := [][4]byte{
		{'M', 'M', 'S', 'S'},
		{'S', 'S', 'M', 'M'},
		{'M', 'S', 'M', 'S'},
		{'S', 'M', 'S', 'M'},
	}

	for _, config := range configs {
		if input[x-1][y-1] == config[0] && input[x-1][y+1] == config[1] &&
			input[x+1][y-1] == config[2] && input[x+1][y+1] == config[3] {
			c++
		}
	}

	return c
}
