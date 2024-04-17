package main

import (
	stringutil "common/Go_revision/utility_functions/stringUtil"
	"fmt"
)

func main() {

	s := "Hello World"
	reversed := stringutil.Reverse(s)
	fmt.Println(reversed)

	fmt.Println(stringutil.IsPalindrome("radar"))
	fmt.Println(stringutil.IsPalindrome("hello"))
}