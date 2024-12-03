package main

func main() {
	ls := parse("input.txt")
	// add mul
	println(add(ls, false))
	println(add(ls, true))

}
