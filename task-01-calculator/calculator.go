package main

import (
	"fmt"
)

func Vbr() {
	var x, y int
	var c string

	fmt.Println("Введите два числа и выберите операцию: ")
	fmt.Scan(&x)
	fmt.Scan(&y)
	fmt.Scan(&c)

	switch c {
	case "+":
		fmt.Println(x, "+", y, "=", x+y)
	case "-":
		fmt.Println(x, "-", y, "=", x-y)
	case "*":
		fmt.Println(x, "*", y, "=", x*y)
	case "/":
		if y != 0 {
			fmt.Println(x, "/", y, "=", x/y)
		} else {
			fmt.Print("нельзя делить на ноль!")
		}
	}
}
