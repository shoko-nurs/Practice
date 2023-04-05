package main

import "fmt"

func main() {

	a := Add(10, 20)

	fmt.Println(a)
	// This is from main line
}

func Add(x, y int) (result int) {

	result = x + y
	return
}

func AddFromFix1() {
	// Removed prvious results and added these new
}
