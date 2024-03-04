package main

import "fmt"

func main() {
	// ex1 array
    // ns := [...]string{
    //     "Almond Biscotti Café", "Banana Pudding", "Balsamic Strawberry (GF)", "Bittersweet Chocolate (GF)",
    //     "Blueberry Pancake (GF)", "Bourbon Turtle (GF)", "Browned Butter Cookie Dough", "Chocolate Covered Black Cherry (GF)",
    //     "Chocolate Fudge Brownie", "Chocolate Peanut Butter (GF)", "Coffee (GF)", "Cookies & Cream", "Fresh Basil (GF)",
    //     "Garden Mint Chip (GF)", "Lavender Lemon Honey (GF)", "Lemon Bar", "Madagascar Vanilla (GF)", "Matcha (GF)",
    //     "Midnight Chocolate Sorbet (GF, V)", "Non-Dairy Chocolate Peanut Butter (GF, V)", "Non-Dairy Coconut Matcha (GF, V)",
    //     "Non-Dairy Sunbutter Cinnamon (GF, V)", "Orange Cream (GF)", "Peanut Butter Cookie Dough", "Raspberry Sorbet (GF, V)",
    //     "Salty Caramel (GF)", "Slate Slate Different", "Strawberry Lemonade Sorbet (GF, V)", "Vanilla Caramel Blondie",
    //     "Vietnamese Cinnamon (GF)", "Wolverine Tracks (GF)",
    // }
    // fmt.Printf("Length of array: %d\n", len(ns))
    // fmt.Printf("Array: %#v\n", ns)
	// fmt.Printf("%T\n",ns)

	//ex2 slice
	// ns:=[]string{"what","is","your","problem"}
	// fmt.Printf("length of the slice: %d\n",len(ns))
	// fmt.Println("Slice:",ns)
	// for index,value:=range ns{
	// 	fmt.Printf("Index:%d, value:%s\n",index,value)
	// }

	// ex3 slice index
	// ns := []string{"Banana Pudding", "Balsamic Strawberry (GF)", "Bittersweet Chocolate (GF)",
	// "Blueberry Pancake (GF)", "Bourbon Turtle (GF)", "Browned Butter"}

    // fmt.Println(ns)
	// for _, value := range ns {
	// 	fmt.Printf("\n%v\n", value)
	// }
	  
	// ex4 slice append
	// ns := []string{"Banana Pudding", "Balsamic Strawberry (GF)", "Bittersweet Chocolate (GF)",
	// "Blueberry Pancake (GF)", "Bourbon Turtle (GF)", "Browned Butter"}

    // fmt.Println(ns)
	// ns=append(ns,"blah","LoL")
	// for _, value := range ns {
	// 	fmt.Printf("\n%v\n", value)
	// }

	// ex5 slicing a slice
	// xi:=[]int{1,2,3,4,5,6}
	// fmt.Printf("%#v\n",xi)
	// fmt.Printf("%#v\n",xi[0:3])
	// fmt.Printf("%#v\n",xi[:5])
	// fmt.Printf("%#v\n",xi[5:])
	// fmt.Printf("%#v\n",xi[:])

	//ex6 deleting a slice
	// xi:=[]int{0,1,2,3,4,5,6,7,8,9}
	// fmt.Printf("%#v\n",xi)
	// xi = append(xi[:4], xi[5:]...)
    // fmt.Printf("%#v\n", xi)

	// ex7 slice 'make'
	// xi:=[]int{0,1,2,3,4,5,6,7,8,9}
	// fmt.Printf("%#v\n",xi)

	// si:=make([]int,0,10)
	// fmt.Println(si)
	// fmt.Println(len(si))
	// fmt.Println(cap(si))
	// si=append(si,0,2,34,55,66)
	// fmt.Println(si)

	//ex8 multidimensionl slice
	// jb:=[]string{"jim","james","lalala"}
	// jm:=[]string{"peter","meg","chris","lois"}
	// fmt.Println(jb)
	// fmt.Println(jm)

	// mix:=[][]string{jb,jm}
	// fmt.Println(mix)

	// ex9 internal & underlying slice
	// a:=[]int{2,3,45,98,66,99}
	// b:=a
	// b:=make([]int,6)
	// copy(b,a)

	// fmt.Println(a)
	// fmt.Println(b)

	// a[0]=9
	// fmt.Println(a)
	// fmt.Println(b)

	// ex10
	// n:=[5]int{}
	// for i:=0;i<5;i++{
	// 	n[i]=i
	// }
	// for i,v:=range n{
	// 	fmt.Printf("Index:%d, value:%d, type:%T\n",i,v,v)
	// }

	// ex11
	// n:=[]int{42,43,44,45,46,47,48,49,50,51}
	// for i,v:=range n{
	// 	fmt.Printf("Index:%d, value:%d, type:%T\n",i,v,v)
	// }

	// ex12
	// n:=[]int{42,43,44,45,46,47,48,49,50,51}
	// a:=n[0:5]
	// b:=n[5:]
	// c:=n[2:7]
	// d:=n[1:6]
	// fmt.Println(n)
	// fmt.Println(a)
	// fmt.Println(b)
	// fmt.Println(c)
	// fmt.Println(d)

	// ex13
	// n:=[]int{42,43,44,45,46,47,48,49,50,51}
	// n=append(n,52)
	// fmt.Println(n)
	// y:=[]int{53,54,55}
	// n=append(n,y...)
	// fmt.Println(n)

	// ex14
	// n:=[]int{42,43,44,45,46,47,48,49,50,51}
	// n= append(n[:3],n[6:]...)
	// fmt.Println(n)

	// ex15
	// s := make([]string, 0)
    // s = append(s, `Alabama`, `Alaska`, `Arizona`, `Arkansas`, `California`, `Colorado`, `Connecticut`,
    //     `Delaware`, `Florida`, `Georgia`, `Hawaii`, `Idaho`, `Illinois`, `Indiana`, `Iowa`, `Kansas`,
    //     `Kentucky`, `Louisiana`, `Maine`, `Maryland`, `Massachusetts`, `Michigan`, `Minnesota`,
    //     `Mississippi`, `Missouri`, `Montana`, `Nebraska`, `Nevada`, `New Hampshire`, `New Jersey`,
    //     `New Mexico`, `New York`, `North Carolina`, `North Dakota`, `Ohio`, `Oklahoma`, `Oregon`,
    //     `Pennsylvania`, `Rhode Island`, `South Carolina`, `South Dakota`, `Tennessee`, `Texas`,
    //     `Utah`, `Vermont`, `Virginia`, `Washington`, `West Virginia`, `Wisconsin`, `Wyoming`,
    // )
    // fmt.Println("Length of slice:", len(s))
    // fmt.Println("Capacity of slice:", cap(s))

    // for i, state := range s {
    //     fmt.Printf("%d: %s\n", i, state)
    // }

	// ex16
	y:=[]string{"ye","yezzy","west"}
	z:=[]string{"drizzy","glizzy","marvin"}
	fmt.Println(y)
	fmt.Println(z)
	yz:=[][]string{y,z}
	fmt.Println(yz)
	for i,v:=range yz{
		fmt.Printf("index:%d\t value:%s\n",i,v)
		for a,b := range v{
			fmt.Println(a,b)
		}
	}

}
