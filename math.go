package main

import "fmt"

func main() {
	fmt.Println(soma(114, 10))
}

func soma(a int, b int) int {
	return a + b
}

func sub(a int, b int) int {
	return a - b
}

func sumX(a int, b int) int {
	return a + b + a
}

func times(a int, b int) int {
	return a * b
}

func sum(a int, b int) int {
	return a + b + a
}
