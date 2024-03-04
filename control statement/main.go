package main

import (
    "fmt"
    // "math/rand"
)
//ex3
// func init(){
// 	fmt.Println("this works before main")
// }
func main() {
	// ex1
    // x := rand.Intn(251)
    // fmt.Printf("Value of x: %d\n", x)
    // if x <= 100 {
    //     fmt.Println("x is between 0 and 100")
    // } else if x <= 200 {
    //     fmt.Println("x is between 101 and 200")
    // } else {
    //     fmt.Println("x is between 201 and 250")
    // }


	//ex2
	// switch{
	// case  x <= 100 :
	// 	fmt.Println("x is between 0 and 100")
	// case x <= 200:
	// 	fmt.Println("x is between 101 and 200")
	// case x<=250:
	// 	fmt.Println("x is between 200 and 250")
	// default:
	// 	fmt.Println("blah")
	// }

	// ex4
	// x:=rand.Intn(10)
	// y:=rand.Intn(8)
	// fmt.Printf("Value of x: %d\n", x)
	// fmt.Printf("Value of y: %d\n", y)
	// if(x<4 && y<4){
	// 	fmt.Println("x and y are both less than 4")
	// }else if(x>6 && y>6){
	// 	fmt.Println("x and y are both greater than 6")
	// }else if(x>=4 && x<=6){
	// 	fmt.Println(" x is greater than or equal to 4 and less than or equal to 6")
	// }else if(y!=5){
	// 	fmt.Println("y is not 5")
	// }else{
	// 	fmt.Println("blahhhh")
	// }

	// ex5
	// switch{
	// case x<4 && y<4:
	// 	fmt.Println("x and y are both less than 4")
	// case x>6 && y>6:
	// 	fmt.Println("x and y are both less than 4")
	// case x>=4 && x<=6:
	// 	fmt.Println(" x is greater than or equal to 4 and less than or equal to 6")
	// case y!=5:
	// 	fmt.Println("y is not 5")
	// default:
	// 	fmt.Println("blahhhh")
	// }

	// ex6
	// x := 100
    // for i := 0; i < x; i++ {
    //     fmt.Println(i)
    // }

	// for i := 0; i < 100; i++ {
	// 	x:=rand.Intn(10)
	// 	y:=rand.Intn(8)
	// 	fmt.Printf("Value of x: %d\n", x)
	// 	fmt.Printf("Value of y: %d\n", y)
	// 	switch{
	// 		case x<4 && y<4:
	// 			fmt.Println("x and y are both less than 4")
	// 		case x>6 && y>6:
	// 			fmt.Println("x and y are both less than 4")
	// 		case x>=4 && x<=6:
	// 			fmt.Println(" x is greater than or equal to 4 and less than or equal to 6")
	// 		case y!=5:
	// 			fmt.Println("y is not 5")
	// 		default:
	// 			fmt.Println("blahhhh")
	// 		}
    // }

	// ex7
	// for i:=1;i<43;i++{
	// 	x:=rand.Intn(5)
	// 	fmt.Printf("Iteration %d: ", i)
	// 	switch x {
	// 	case 0:
	// 		fmt.Println("x is 0")
	// 	case 1:
	// 		fmt.Println("x is 1")
	// 	case 2:
	// 		fmt.Println("x is 2")
	// 	case 3:
	// 		fmt.Println("x is 3")
	// 	case 4:
	// 		fmt.Println("x is 4")
	// 	default:
	// 		fmt.Println("x is not in the range [0, 4]")
	// 	}
	// }

	// ex 8
	// x:=10
	// for x>0{
	// 	fmt.Printf("Iteration %d\n",x)
	// 	x--
	// }

	// ex9
	// for i:=1;i<10;i++{
	// 	if i%2==0{
	// 		fmt.Printf("Even number %d\n",i)

	// 	}else{
	// 		fmt.Printf("Odd number %d\n",i)
	// 		break
	// 	}
	// }

	// ex10
	// for i:=1;i<10;i++{
	// 		if i%2==0{
	// 			fmt.Printf("Even number %d\n",i)
	
	// 		}else{
	// 			continue
	// 		}
	// 	}

	// ex11
    // for i := 1; i <= 5; i++ {
    //     fmt.Printf("Outer loop iteration %d\n", i)
    //     for j := 1; j <= 5; j++ {
    //         fmt.Printf("\tInner loop iteration %d\n", j)
    //     }
    // }

	// ex12
	// xi := []int{42, 43, 44, 45, 46, 47}
	// for index,value:=range xi{
	// 	fmt.Printf("Index:%d\tvalue:%d\n",index,value)
	// }

	// ex13
	// m := map[string]int{
	// 	"James": 42,
	// 	"Moneypenny": 32,
	// 	}
	// for i,v:=range m{
	// 	fmt.Printf("Index:%s\t Value:%d\n",i,v)
	// }

	//ex14
	// for i:=0;i<101;i++{
	// 	if x:=rand.Intn(5);x==3{
	// 		fmt.Printf("iteration:%d x is 3\n",i)
	// 	}
	// }

	// ex15
	// fmt.Println(true && true)
	// fmt.Println(true && false)
	//  fmt.Println(true || true)
	//  fmt.Println(true || false)
	// fmt.Println(!true)
}



