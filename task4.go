package main

import "fmt"

func main() {
	num := 123123

	
	firstPart := num / 1000 
	lastPart := num % 1000  

	
	sum1 := (firstPart / 100) + ((firstPart / 10) % 10) + (firstPart % 10)


	sum2 := (lastPart / 100) + ((lastPart / 10) % 10) + (lastPart % 10)


	if sum1 > sum2 && sum2 > sum1{
		fmt.Println("NO")

	} else {
		fmt.Println("YES")
	}
}
