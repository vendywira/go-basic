package main

import "fmt"

func main() {
	/*
		- arithmetic operator
	 */

	var add int = 3 + 4
	var substract int = 3 - 4
	var multy int = 3 * 4
	var division float32 = 4 / 3
	var modulo int = 4 % 3
	fmt.Println(add, substract, multy, division, modulo)

	var calculate float32 = (((2 + 6) % 3) * 4 - 2) / 3
	fmt.Println(calculate)

	/*
		comparison operator -> will use for get boolean value
	 */

	var equal bool = "a" == "a"
	fmt.Println(equal)

	var notEqual bool = "a" != "b"
	fmt.Println(notEqual)

	greaterThan := 3 > 4
	fmt.Println(greaterThan)

	lessThan := 3 < 4
	fmt.Println(lessThan)

	greaterThanEqual := 3 >= 4
	fmt.Println(greaterThanEqual)

	lessThanEqual := 3 <= 4
	fmt.Println(lessThanEqual)

	/*
		Logic operator
	 */

	 and := equal && notEqual
	 fmt.Println("and condition \t: \t", and)

	 or := lessThan || greaterThan
	 fmt.Println("or condition \t: \t", or)

	 neagation := !equal
	 fmt.Println("negation \t\t: \t", neagation)
}
