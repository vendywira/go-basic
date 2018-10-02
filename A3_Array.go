package main

import "fmt"

func main() {

	// declare array
	var names [4]string
	names[0] = "I"
	names[1] = "Wayan"
	names[2] = "Vendy"
	names[3] = "Wiranatha"
	fmt.Println(names)

	var array = [5]int{1, 2, 3, 4, 5}
	fmt.Println("index 0: ", array[0], "index 3: ", array[3])

	// set array
	array[1] = 10
	fmt.Println("index 1 : ", array[1])

	// declare array fruits in vertical manner. last element also should be used ","
	fruits := [5]string{
		"banana",
		"apple",
		"grape",
		"water melon",
		"melon",
	}
	fmt.Println(fruits)

	// declare array auto detect length of it
	numbers := [...]uint8{1, 2, 3, 4}
	fmt.Println(numbers, " with length ", len(numbers))

	// declare use keyword make
	var animals = make([]string, 2)
	animals[0] = "cat"
	animals[1] = "dog"
	fmt.Println(animals)
}
