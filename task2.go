package main

import "fmt"

func main() {
	num := 385

	c1 := num / 100
	c2 := (num / 10) % 10
	c3 := num % 10

	if c1 != c2 && c1 != c3 && c2 != c3 {
		fmt.Println("Yes")
	} else {
		fmt.Println("NO")

	}
}
