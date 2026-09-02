package main

import "fmt"

func main() {
	num := 5678
	var res int

	for num >= 10 {
		res = num / 10
		num = res
	}
	fmt.Println(res)
}
