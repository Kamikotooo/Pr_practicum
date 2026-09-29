package main

import "fmt"

func main() {
	var a int = 10
	var a8 int8 = 8
	var a16 int16 = 16
	var a32 int32 = 32
	var a64 int64 = 64
	var str string = "Просто строка"
	var f32 float32 = 3.14
	var f64 float64 = 2.754
	var b bool = true
	var bt byte = 'A'
	var r rune = 'A'

	var a_v = 20
	var f_v = 3.33
	var str_v = "ЭЭЭЭЭЭЭ"
	var b_v = true

	a_pr := 20
	f_pr := 4.44
	str_pr := "символ"
	b_pr := true

	var x int = 10
	var y int = 20
	var x_1 int = 100
	var y_1 int = 200
	fmt.Printf("\n", x, y)
	x, y = y, x
	fmt.Printf("\n", x, y)

	fmt.Printf("\n", x_1, y_1)
	temp := x_1
	x_1 = y_1
	y_1 = temp
	fmt.Printf("\n", x_1, y_1)

	fmt.Println("\n _____var с Типом_________ \n")
	fmt.Printf("\n%T", a, a)
	fmt.Printf("\n%T", a8, a8)
	fmt.Printf("\n%T", a16, a16)
	fmt.Printf("\n%T", a32, a32)
	fmt.Printf("\n%T", a64, a64)
	fmt.Printf("\n%T", str, str)
	fmt.Printf("\n%T", f32, f32)
	fmt.Printf("\n%T", f64, f64)
	fmt.Printf("\n%T", b, b)
	fmt.Printf("\n%T", bt, bt)
	fmt.Printf("\n%T", r, r)

	fmt.Println("\n _____var без Типа_________ \n")
	fmt.Printf("\n%T", a_v, a_v)
	fmt.Printf("\n%T", f_v, f_v)
	fmt.Printf("\n%T", str_v, str_v)
	fmt.Printf("\n%T", b_v, b_v)

	fmt.Println("\n _____ Присваивание :=_________ \n")
	fmt.Printf("\n%T", a_pr, a_pr)
	fmt.Printf("\n%T", f_pr, f_pr)
	fmt.Printf("\n%T", str_pr, str_pr)
	fmt.Printf("\n%T", b_pr, b_pr)

}
