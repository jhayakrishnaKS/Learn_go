package main

import "fmt"

func main() {
	// array
	// arr := [...]string{
	// 	"Almond Biscotti Café", "Banana Pudding ", "Balsamic Strawberry (GF)",
	// 	"BittersweetChocolate (GF)", "Blueberry Pancake (GF)", "Bourbon Turtle (GF)",
	// 	"Browned Butter Cookie Dough", "Chocolate Covered Black Cherry (GF)",
	// 	"Chocolate Fudge Brownie", "Chocolate Peanut Butter (GF)", "Coffee (GF)",
	// 	"Cookies & Cream", "Fresh Basil (GF)", "Garden Mint Chip (GF)",
	// 	"Lavender Lemon Honey (GF)", "Lemon Bar", "Madagascar Vanilla (GF)",
	// 	"Matcha (GF)", "Midnight Chocolate Sorbet (GF, V)",
	// 	"Non-Dairy Chocolate Peanut Butter (GF, V)", "Non-Dairy Coconut Matcha (GF, V)",
	// 	"Non-Dairy Sunbutter Cinnamon (GF, V)", "Orange Cream (GF)",
	// 	"Peanut Butter Cookie Dough", "Raspberry Sorbet (GF, V)", "Salty Caramel (GF)",
	// 	"Slate Slate Different", "Strawberry Lemonade Sorbet (GF, V)",
	// 	"Vanilla Caramel Blondie", "Vietnamese Cinnamon (GF)", "Wolverine Tracks (GF)",
	// }
	// fmt.Println(len(arr))
	// fmt.Println(arr)

	//slice
	// slic := []string{
	// 	"Almond Biscotti Café", "Banana Pudding ", "Balsamic Strawberry (GF)",
	// 	"BittersweetChocolate (GF)", "Blueberry Pancake (GF)", "Bourbon Turtle (GF)",
	// 	"Browned Butter Cookie Dough", "Chocolate Covered Black Cherry (GF)",
	// 	"Chocolate Fudge Brownie", "Chocolate Peanut Butter (GF)", "Coffee (GF)",
	// 	"Cookies & Cream", "Fresh Basil (GF)", "Garden Mint Chip (GF)",
	// 	"Lavender Lemon Honey (GF)", "Lemon Bar", "Madagascar Vanilla (GF)",
	// 	"Matcha (GF)", "Midnight Chocolate Sorbet (GF, V)",
	// 	"Non-Dairy Chocolate Peanut Butter (GF, V)", "Non-Dairy Coconut Matcha (GF, V)",
	// 	"Non-Dairy Sunbutter Cinnamon (GF, V)", "Orange Cream (GF)",
	// 	"Peanut Butter Cookie Dough", "Raspberry Sorbet (GF, V)", "Salty Caramel (GF)",
	// 	"Slate Slate Different", "Strawberry Lemonade Sorbet (GF, V)",
	// 	"Vanilla Caramel Blondie", "Vietnamese Cinnamon (GF)", "Wolverine Tracks (GF)",
	// }

	// fmt.Println(len(slic))
	// fmt.Printf("%T\n",slic)

	// for i,v:=range slic{
	// 	fmt.Printf("index:%d--value:%s\n",i,v)
	// }
	//appending
	// slic=append(slic, "jkay:blueberry")
	// fmt.Print(slic,"\n")

	//slice
	// slic=slic[0:5]
	// fmt.Print(slic,"\n")

	sli := []int{1, 2, 3, 4, 5}
	modifiedSlice := append(sli[:1], sli[2:]...)

	fmt.Println(modifiedSlice)

	m := []string{"James", "Bond", "Shaken, not stirred"}
	f := []string{"Miss", "Moneypenny", "I'm 008."}

	mf := [][]string{m, f}

	for i, v := range mf {
		fmt.Println(i, v)
		for a, b := range v {
			fmt.Println(a, b)
		}
	}

}
