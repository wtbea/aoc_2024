package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

const (
	UP = iota
	RIGHT
	DOWN
	LEFT
)

func p1(ls []string) int {
	var grid [][]byte
	for _, line := range ls {
		lineCopy := make([]byte, len(line))
		copy(lineCopy, line)
		grid = append(grid, lineCopy)
	}

	p1 := 1
	timeSinceLastNew := 0
	dir := -1
	var x, y int

found:
	for i, row := range grid {
		for j, c := range row {
			switch c {
			case '^':
				dir = UP
			case '>':
				dir = RIGHT
			case 'V', 'v':
				dir = DOWN
			case '<':
				dir = LEFT
			default:
				continue
			}

			if dir != -1 {
				x = i
				y = j
				grid[x][y] = 'X'
				break found
			}
		}
	}

	for timeSinceLastNew < p1 {
		timeSinceLastNew++

		switch dir {
		case UP:
			x--
		case DOWN:
			x++
		case LEFT:
			y--
		case RIGHT:
			y++
		}

		if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[0]) {
			break
		}

		if grid[x][y] == '#' {
			switch dir {
			case UP:
				x++
			case DOWN:
				x--
			case LEFT:
				y++
			case RIGHT:
				y--
			}
			dir = (dir + 1) % 4
			continue
		}
		if grid[x][y] != 'X' {
			grid[x][y] = 'X'
			p1++
			timeSinceLastNew = 0
		}
	}
	return p1
}

func move(grid [][]byte, x *int, y *int, dir *int) error {
	for i := 0; i < 4; i++ {
		switch *dir {
		case UP:
			*x--
		case DOWN:
			*x++
		case LEFT:
			*y--
		case RIGHT:
			*y++
		}

		if *x < 0 || *x >= len(grid) || *y < 0 || *y >= len(grid[*x]) {
			return fmt.Errorf("out of bounds")
		}

		if grid[*x][*y] == '#' {
			switch *dir {
			case UP:
				*x++
			case DOWN:
				*x--
			case LEFT:
				*y++
			case RIGHT:
				*y--
			}
			*dir = (*dir + 1) % 4
			continue
		} else {
			break
		}
	}
	return nil
}

func checkIfLoop(grid [][]byte, x, y, dir int) bool {
	prevVisited := make([][]int, len(grid))
	for i := range prevVisited {
		prevVisited[i] = make([]int, len(grid[i]))
	}

	countCurr := 1
	xCopy := x
	yCopy := y
	currentDir := dir

	for prevVisited[xCopy][yCopy] != countCurr {
		prevVisited[xCopy][yCopy] = countCurr
		prevX := xCopy
		prevY := yCopy

		err := move(grid, &xCopy, &yCopy, &currentDir)
		if err != nil {
			return false
		}

		if xCopy == prevX && yCopy == prevY {
			return true
		}

		if grid[xCopy][yCopy] != 'X' {
			grid[xCopy][yCopy] = 'X'
			countCurr++
		}
	}

	return true
}

func checkRow(grid [][]byte, x, y, rowNum, dir int, out *int32, counter *int32, wg *sync.WaitGroup) {
	defer wg.Done()
	calcVal := 0
	for i := 0; i < len(grid[rowNum]); i++ {
		copyGrid := make([][]byte, len(grid))
		for j := range grid {
			copyGrid[j] = make([]byte, len(grid[j]))
			copy(copyGrid[j], grid[j])
		}

		copyGrid[rowNum][i] = '#'
		if checkIfLoop(copyGrid, x, y, dir) {
			calcVal++
		}
	}

	atomic.AddInt32(out, int32(calcVal))
	atomic.AddInt32(counter, 1)
}

func path(ls []string, v2 bool) int {
	if v2 {
		return p2(ls)
	}
	return p1(ls)
}

func p2(ls []string) int {
	var grid [][]byte
	for _, line := range ls {
		lineCopy := make([]byte, len(line))
		copy(lineCopy, line)
		grid = append(grid, lineCopy)
	}

	dir := -1
	var x, y int
	found := false
	for i, row := range grid {
		for j, c := range row {
			switch c {
			case '^':
				dir = UP
			case '>':
				dir = RIGHT
			case 'V', 'v':
				dir = DOWN
			case '<':
				dir = LEFT
			default:
				continue
			}

			if dir != -1 {
				x = i
				y = j
				grid[x][y] = 'X'
				found = true
				break
			}
		}
		if found {
			break
		}
	}

	var p2 int32 = 0
	var counter int32 = 0
	var wg sync.WaitGroup

	for i := 0; i < len(grid); i++ {
		wg.Add(1)
		go checkRow(grid, x, y, i, dir, &p2, &counter, &wg)
	}

	wg.Wait()

	return int(p2)
}
