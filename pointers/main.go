package main

import (
	"fmt"

	// "golang.org/x/text/width"
)

// func inDelta(n int) int {
// 	return 43
// }

// func inDeltaSlice(n []int) {
// 	n[0] = 99
// }

// func inDeltaMap(n map[string]int, s string) {
// 	n[s] = 90
// }

type dog struct{
	first string
}

func(d dog)walk(){
	fmt.Println("may name is",d.first,"and i'm walking")
}
func(d *dog)run(){
	d.first="rover"
	fmt.Println("may name is",d.first,"and i'm running")
}
type youngIn interface{
	walk()
	run()
}
func youngRun(y youngIn){
	y.run()
}

func main() {
	// x := 42
	// fmt.Println(x)
	// fmt.Println(&x)

	// y:=&x
	// fmt.Printf("%v \t %T\n",y,y)
	// fmt.Println(y)

	// fmt.Println(*&x)
	// *y=43
	// fmt.Println(x)
	// fmt.Println(*y)

	// a:=42
	// fmt.Println(a)
	// fmt.Println(inDelta(a))
	// fmt.Println(a)

	// xi:=[]int{1,2,34,45}
	// fmt.Println(xi)
	// inDeltaSlice(xi)
	// fmt.Println(xi)

	// m:=make(map[string]int)
	// m["jkay"]=43
	// fmt.Println(m)
	// inDeltaMap(m,"james")
	// fmt.Println(m["james"])

	//value semantics
	// b := 1
	// fmt.Println(b)
	// fmt.Println(addOne(b))
	// fmt.Println(b)

	//pointer semantics
	// a:=1
	// fmt.Println("pointer semantics\n",a)
	// addOneP(&a)
	// fmt.Println(a)

	d1:=dog{"terry"}
	d1.walk()
	d1.run() 


	d2:=&dog{"nicey"}
	d2.walk()
	d2.run()
	youngRun(d2)

}

//value semantics
// func addOne(b int) int {
// 	return b + 1
// }

//pointer semantics
// func addOneP(b *int) {
// 	 *b+= 1
// }
