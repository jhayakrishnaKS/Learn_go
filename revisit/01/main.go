package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// you can range over channels
	// ch := make(chan int)
	// go func() {
	// 	ch <- 1
	// 	ch <- 2
	// 	close(ch)
	// }()
	// for value := range ch {
	// 	fmt.Println(value)
	// }

	// m:=make(map[string]int)
	// m["ten"]=10
	// m["eleven"]=11
	// fmt.Println(m)

	// x := rand.Intn(251)
	// if x==0 && x<=100{
	// 	fmt.Println("print between 0 and 100",x)
	// }else if x<=200{
	// 	fmt.Println("the value is between 101 and 200",x)
	// }else{
	// 	fmt.Println("the value is between 201 and 250",x)
	// }

	// switch {
	// case x==0 && x<=100:
	// 	fmt.Println("print between 0 and 100",x)
	// case x<=200:
	// 	fmt.Println("the value is between 101 and 200",x)
	// default:
	// 	fmt.Println("the value is between 201 and 250",x)
	// }

	for i := 0; i < 101; i++ {
		if x := rand.Intn(5); x == 1 {
			fmt.Printf("iteration:%d x is 1\n", i)
		}
	}
}
