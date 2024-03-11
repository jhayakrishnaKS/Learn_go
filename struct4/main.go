package main

import (
	"fmt"
)

// type person struct {
//     first string
//     last  string
//     age   int
// }

// type agent struct{
// 	person
// 	lkt bool
// }

// type person struct{
// 	firstName string
// 	lastName string
// 	iceCream []string
// }

// type engine struct{
// 	electric bool
// }
// type vechicle struct{
// 	engine
// 	make string
// 	model string
//     doors int
//     color string
// }

func main() {
	// p1 := person{
	//     first: "james",
	//     last:  "harpert",
	//     age:   28,
	// }
	// p2 := person{
	//     first: "dwight",
	//     last:  "schrute",
	//     age:   29,
	// }
	// fmt.Println(p1.first)
	// fmt.Println(p2.last)

	//embedded struct
	// p1 := agent{
	// 	person:person{
	// 	first: "james",
	//     last:  "harpert",
	//     age:   28,
	// 	},
	// 	lkt:true,
	// }
	// fmt.Println(p1)
	// fmt.Println(p1.first,p1.last,p1.age,p1.lkt)

	//anonymous struct
	// p1:=struct{
	// 	first string
	//     last  string
	//     age   int
	// }{
	// 	first: "james",
	//     last:  "harpert",
	//     age:   28,
	// }
	// fmt.Printf("%T",p1)

	// ex1
	// p1:=person{
	// 	firstName:"krish",
	// 	lastName:"Jk",
	// 	iceCream:[]string{"chocolate","sorbet","vanilla"},
	// }
	// p2:=person{
	// 	firstName:"zunade",
	// 	lastName:"james",
	// 	iceCream:[]string{"strawberry","desire","lychee"},
	// }

	// fmt.Println(p1)
	// fmt.Println(p2)

	// fmt.Println(p1.iceCream)
	// fmt.Println(p2.iceCream)

	// for _,v:=range p1.iceCream{
	// 	fmt.Println(p1.firstName,"'s favorite",v)
	// }

	// for _,v:=range p2.iceCream{
	// 	fmt.Println(p2.firstName,"'s favorite",v)
	// }

	// ex2
	// p1:=person{
	// 	firstName:"krish",
	// 	lastName:"Jk",
	// 	iceCream:[]string{"chocolate","sorbet","vanilla"},
	// }
	// p2:=person{
	// 	firstName:"zunade",
	// 	lastName:"james",
	// 	iceCream:[]string{"strawberry","desire","lychee"},
	// }
	// m:=map[string]person{
	// 	p1.lastName:p1,
	// 	p2.lastName:p2,
	// }
	// for _,v:=range m{
	// 	fmt.Println(v)
	// 	for _,v2:=range v.iceCream{
	// 		fmt.Println(v.firstName,v.lastName,v2)
	// 	}
	// }

	// ex3
	// v1:=vechicle{
	// 	engine:engine{
	// 		electric:false,
	// 	},
	// 	make:"indian",
	// 	model:"v1",
	// 	doors:4,
	// 	color:"jet black",
	// }
	// v2:=vechicle{
	// 	engine:engine{
	// 		electric:true,
	// 	},
	// 	make:"american",
	// 	model:"cyberTruck",
	// 	doors:2,
	// 	color:"astro white",
	// }
	// fmt.Println(v1.engine.electric)
	// fmt.Println(v2.make)

	// ex4
	p1 := struct {
		first     string
		friends   map[string]int
		favdrinks []string
	}{
		first: "jkay",
		friends: map[string]int{
			"rave": 1,
			"zun":  2,
		},
		favdrinks: []string{"sx", "cranberry", "zx"},
	}

	for k, v := range p1.friends {
		fmt.Println(p1.first, "-friend-", k, ":", v)
	}

	for _, v := range p1.favdrinks {
		fmt.Println(p1.first, "-drinks-", v)
	}

}
