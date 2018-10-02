package main

import "fmt"

func main() {

	/*
		- data type numeric and not decimal last digit define large memory we need
		- uint value start from 0 and can't set negative value
		- int can set with negative value
		- byte as same as uint8
		- uint will auto sign to uint32 or uint64 base on data assigned
		- rune as same as int32
		- int will auto sign to int32 or int64
	 */
	var uInteger_8 uint8 = 8
	var uInteger_16 uint16 = 16
	var uInteger_32 uint32 = 32
	var uInteger_64 uint64 = 64
	var byte_uint8  byte = 8
	var uInteger_32_64 uint = 100000000
	fmt.Println(uInteger_8, uInteger_16, uInteger_32, uInteger_64, byte_uint8, uInteger_32_64)

	var int_8 int8 = -8
	var int_16 int16 = 16
	var int_32 int32 = -32
	var int_64 int64 = 50
	var int_32_64 int = -297384474676
	var rune_int_32 rune = 345
	fmt.Println(int_8, int_16, int_32, int_64, int_32_64, rune_int_32)


	/*
		- data type numeric decimal last digit define large memory we need
		- on golang for decimal only have 2 data type float32 and float64
		- if you java developer before float64 like double on java
	 */
	var float_32 float32 = 2.78
	var float_64 float64 = -32.00897837403
	fmt.Println(float_32, float_64)

	/*
		- data type bool (Boolean) just have 2 value true or false
	 */
	var trueValue bool = true
	var falseValue bool = false
	fmt.Println(trueValue, falseValue)

	// syntax for declare variable on golang used data type
	var firstWord string = "first word"

	var secondWord string
	secondWord = "second word"

	fmt.Printf("these are %s, %s \n", firstWord, secondWord)

	// syntax for declare variable on golang without data type
	var thirdWord = "third word"
	fmt.Println(thirdWord)

	// declare multiple variable
	var first, second, third, fourth = 1, 2, 3, 4
	fmt.Println(first, second, third, fourth)

	// declare variable without keyword var
	five := 5
	fmt.Println(five)

	// declare variable will not use
	var name, _ = "vendy", "not used var"
	fmt.Println(name)
	fmt.Println("memory location : ", &name)

	// variable pointer for know where location on memory
	pointerVar := new(string)
	fmt.Println(pointerVar)
	fmt.Println(*pointerVar)

	// like variable but value of constanta can't change
	const constanta string = "this is constanta"
	// will get error when change value on constanta
	//constanta = "hello"
	fmt.Println(constanta)

}
