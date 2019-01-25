package main

import "fmt"

func main() {
	/*
	 *	slice is array with reference
	 * 	declare slice use [] without length [5] or verdict [...]
	 */
	var fruits = []string{"apple", "grape", "banana", "water melon", "melon"}
	fmt.Println(fruits[0])

	// to get value fruits[0:3] 0 = start index 3 = get value till index previous 3 or get till index 2
	fmt.Println(fruits[0:3])

	fruits2 := fruits
	fmt.Println(fruits2)

	/*
	 *	don't forget slice is array with reference
	 *	when array reference changed all slice reference from it will changed too.
	 */
	fruits[0] = "orange"
	fmt.Println(fruits2)

	fruits3 := fruits2[0:3]
	fmt.Println(fruits3)

	fruits4 := fruits3[0:2]
	fmt.Println(fruits4)

	// nested reference
	fruits[0] = "apple"
	fmt.Println(fruits2)
	fmt.Println(fruits3)
	fmt.Println(fruits4)
}
