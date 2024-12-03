package main

func main() {
	// distance
	left, right := parse("input.txt")
	println(distance(left, right))

	// similarity
	println(similarity(left, right))
}
