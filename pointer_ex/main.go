package main

import "fmt"

//ex2
// var (
// 	a, b, c *string
// 	d       *int
// )

// func init() {
// 	p := "Drop by drop, the bucket gets filled."
// 	q := "Persistently, patiently, you are bound to succeed."
// 	r := "The meaning of life is ..."
// 	n := 42
// 	a = &p
// 	b = &q
// 	c = &r
// 	d = &n
// }

// ex4
type person struct {
	first string
}

func changeName(p person, s string) person {
	p.first = s
	return p
}

func changeNameP(p *person, s string) {
	p.first = s
}

func main() {
	// ex1
	// a := 10
	// fmt.Println(a)
	// fmt.Println(&a)

	// ex2
	// fmt.Printf("%v\t%T\n", a, a)
	// fmt.Printf("%v\t%T\n", b, b)
	// fmt.Printf("%v\t%T\n", c, c)
	// fmt.Printf("%v\t%T\n", d, d)

	// fmt.Println(*a)
	// fmt.Println(*b)
	// fmt.Println(*c)
	// fmt.Println(*d)

	//ex3
	// d1 := dog{"Henry"}
	// youngRun(d1)

	// d2 := dog{"Padget"}
	// youngRun(d2)
	// d2 = d2.changeName("Rover")
	// youngRun(d2)

	// ex4
	p := person{
		first: "krish",
	}
	fmt.Println(p)
	fmt.Println(changeName(p, "jk"))
	fmt.Println(p)

	changeNameP(&p, "jlo")
	fmt.Println(p)

}

//ex3
// type dog struct {
// 	first string
// }

// func (d dog) walk() {
// 	fmt.Println("My name is", d.first, "and I'm walk.")
// }

// func (d dog) run() {
// 	fmt.Println("My name is", d.first, "and I'm running.")
// }

// type youngin interface {
// 	walk()
// 	run()
// }

// func youngRun(y youngin) {
// 	y.walk()
// 	y.run()
// }

// func (d dog) changeName(s string) dog {
// 	d.first = s
// 	return d
// }
