//Single line comment

/*
	multiline comments
	span into more than 1 line.
*/

//Go file starts with package name. main is special package containing main function.
//As a convention, package name is all lowercase with no special character or space.
//Generally, package name is taken as the name of directory.

//For test, package name ends in _test suffix.
package main

import (
	"fmt"
	"learn-go/pack"
)

//Name are snakeCase or PascalCase. snakeCase names are for private constructs which are not exported.
//Be it variable, structure fields, function name.
//PascalCase names are for exportable constructs.

//Entry point.
func main() {
	fmt.Println("Hello world")
	fmt.Println(pack.Sum(3, 5))
	pack.Demo1()

	//Variable declaration. It can be done as

	//Using var and initial value.
	var a = "initial"
	//Printed as unused variables or imports result in compile error.
	fmt.Print(a)

	//Using var and data type
	var b, c int
	b, c = 1, 2
	fmt.Print(b, c)

	//Using initial value
	f := "apple" //f is string.
	fmt.Print(f)

	learnArrays()
	learnSlices()
	learnLoop()
	learnIfElse()
	learnSwitch()
	demo1()
	pack.LearnGoroutine()

}

func learnSlices() {
	// Slices have dynamic size. Arrays and slices each have advantages
	s3 := []int{4, 5, 9} // Compare to a5. No ellipsis here.
	s4 := make([]int, 4) // Allocates slice of 4 ints, initialized to all 0.
	fmt.Print(s4)
	//var d2 [][]float64      // Declaration only, nothing allocated here.
	// Slices have reference semantics.
	s3_cpy := s3                    // Both variables point to the same instance.
	s3_cpy[0] = 0                   // Which means both are updated.
	fmt.Println(s3_cpy[0] == s3[0]) // true
	s := []int{1, 2, 3}             // Result is a slice of length 3.
	s = append(s, 4, 5, 6)          // Added 3 elements. Slice now has length of 6.
	fmt.Print(s)
}

func learnArrays() {
	var a4 [4]int                    // An array of 4 ints, initialized to all 0.
	a5 := [...]int{3, 1, 5, 10, 100} // An array initialized with a fixed size of five
	fmt.Print(a5)
	a4_cpy := a4                    // a4_cpy is a copy of a4, two separate instances.
	a4_cpy[0] = 25                  // Only a4_cpy is changed, a4 stays the same.
	fmt.Println(a4_cpy[0] == a4[0]) // false
}

func learnLoop() {
	for i := 0; i < 4; i++ {
		fmt.Printf("yay")
	}
	/*for {
		fmt.Printf("inflinite loop")
	  }*/
	rvariable := []string{"a", "b", "c"}
	for i, j := range rvariable {
		fmt.Println(i, j) //  0 a
	}

	i := 0
	for i < 3 {
		i += 2
	}

}

func learnIfElse() {
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}

func learnSwitch() {
	var value string = "five"
	switch value {
	case "one":
		fmt.Println("C#")
	case "two", "three":
		fmt.Println("Go")
	case "four", "five", "six":
		fmt.Println("Java")
	}
}

//Function declaration
func test(x int) {
	fmt.Println("value of x:", x)
}

//May return multiple values.
func test1(x, y int) (int, int) {
	return x + y, x * y
}

//Named return argument.
func test2(x, y int) (sum, prod int) {
	sum, prod = x+y, x*y
	return
}

//Function can be passed as argument
func test3(oper func(x, y int) int, o1, o2 int) int {
	return oper(o1, o2)
}

func demo1() {
	fmt.Println(test3(func(x, y int) int { return x + y }, 2, 3))
}
