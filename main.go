package main

import "fmt"

func main() {

	a := Add(10, 20)

	fmt.Println(a)
}

func Add(x, y int) (result int) {

	result = x + y
	return
}

func AddFromFix1() {
	fmt.Println("From fix1 first commit")
	// for commit 2 fix1
}
