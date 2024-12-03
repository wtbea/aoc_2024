package main

func main() {
	// safety
	list := parse("input.txt")
	println(safety(list, false))

	// dampener
	println(safety(list, true))
}
