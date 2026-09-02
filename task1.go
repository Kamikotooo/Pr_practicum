
package main

import "fmt"

func celka(a int) string {
	if a > 0 {
		return "Число положительное"
	} else if a < 0 {
		return "Число отрицательное"
	} else {
		return "Ноль"
	}
}

func main() {
	fmt.Println(celka(5))
}