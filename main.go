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
	// Removed prvious results and added these new
}

func new1() {
	fmt.Println(123)
}

func new2() {
	fmt.Println("Hi")
}
