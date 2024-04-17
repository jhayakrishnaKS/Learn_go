package main

import (
	"fmt"
	// "strings"
	// "math"
	// "time"
	// "log"
	// "math"
	// "golang.org/x/text/width"
	// "os"
)

//ex1
// func bar()(int,string){
// return 42, "bruh"
// }

// func foo()int{
//     return 42
// }

// ex2
func foo(i ...int) int {
	t := 0
	for _, v := range i {
		t += v
	}
	return t
}
func bar(x []int) int {
	t := 0
	for i := 0; i < len(x); i++ {
		t += x[i]
	}
	return t
}

// ex4
// type person struct{
//     first string
//     age int
// }
// func (p person) speak(){
//     fmt.Println("My nickName is:",p.first,"and my age is:",p.age)
// }

// ex5
// type Square struct {
// 	length float64
// 	width  float64
// }
// type Circle struct {
// 	radius float64
// }

// func (s Square) area() float64 {
// 	return s.length * s.width
// }

// func (c Circle) area() float64 {
// 	return math.Pi * math.Pow(c.radius, 2)
// }

// type shape interface {
// 	area() float64
// }

// func info(s shape) float64 {
// 	return s.area()
// }

// ex6
// func doMath(a int, b int, f func(int, int) int) int {
// 	return f(a, b)
// }
// func add(a int, b int) int {
// 	return a + b
// }
// func subtract(a int, b int) int {
// 	return a - b
// }

// ex7
// type User struct {
// 	ID    int
// 	First string
// }
// type MockDatastore struct {
// 	Users map[int]User
// }

// func (md MockDatastore) GetUser(id int) (User, error) {
// 	user, ok := md.Users[id]
// 	if !ok {
// 		return User{}, fmt.Errorf("User %v not found", id)
// 	}
// 	return user, nil
// }

// func (md MockDatastore) SaveUser(u User) error {
// 	_, ok := md.Users[u.ID]
// 	if ok {
// 		return fmt.Errorf("User %v already exists", u.ID)
// 	}
// 	md.Users[u.ID] = u
// 	return nil
// }

// type Datastore interface {
// 	GetUser(id int) (User, error)
// 	SaveUser(u User) error
// }
// type Service struct {
// 	ds Datastore
// }

// func (s Service) GetUser(id int) (User, error) {
// 	return s.ds.GetUser(id)
// }

// func (s Service) SaveUser(u User) error {
// 	return s.ds.SaveUser(u)
// }

// ex9
// func bar() func() int{
// 	return func() int {
// 		return 21
// 	}
//  }

// ex10 callback
// func Square(n int)int{
// 	return n*n
// }
// func printSquare(f func(int) int, num int) string {
//     result := f(num)
//     return fmt.Sprintf("Square of %d is %d", num, result)
// }

// ex11 closure
// func powr(a float64)func() float64{
// 	var c float64
// 	return func() float64 {
// 		c++
// 		return math.Pow(a,c)
// 	}
// }

func main() {
	// Open a file for writing
	// file, err := os.Create("output.txt")
	// if err != nil {
	//     fmt.Println("Error:", err)
	//     return
	// }
	// defer file.Close()

	// Write data to the file using io.Writer interface
	// data := []byte("Hello, world!\n")
	// _, err = file.Write(data)
	// if err != nil {
	//     fmt.Println("Error:", err)
	//     return
	// }

	// fmt.Println("Data has been written to output.txt")

	// ex1
	// fmt.Println(foo())
	// fmt.Println(bar())

	// ex2 vardiadic func
	x := []int{1, 2, 34, 5, 6, 7, 8}
	fmt.Println(foo(x...))
	fmt.Println(bar(x))

	// ex3 defer
	// defer fmt.Print("mfss")
	// fmt.Println("heloo")

	// ex4 method
	// x:=person{
	//     first: "jkay",
	//     age:21,
	// }
	// x.speak()

	// ex5 interfaces
	// s := Square{
	// 	length: 40.1,
	// 	width:  20.5,
	// }
	// c := Circle{
	// 	radius: 5.5,
	// }

	// fmt.Println(info(c))
	// fmt.Println(info(s))

	// ex6 testing
	// fmt.Printf("%T\n", add)
	// fmt.Printf("%T\n", subtract)
	// fmt.Printf("%T\n", doMath)
	// x := doMath(42, 16, add)
	// fmt.Println(x)
	// y := doMath(42, 16, subtract)
	// fmt.Println(y)

	// ex7
	// db := MockDatastore{
	// 	Users: make(map[int]User),
	// }

	// srvc := Service{
	// 	ds: db,
	// }

	// u1 := User{
	// 	ID:    1,
	// 	First: "James",
	// }

	// err := srvc.SaveUser(u1)
	// if err != nil {
	// 	log.Fatalf("error %s", err)
	// }

	// u1Returned, err := srvc.GetUser(u1.ID)
	// if err != nil {
	// 	log.Fatalf("error %s", err)
	// }

	// fmt.Println(u1)
	// fmt.Println(u1Returned)

	// ex8 anonymous func
	// func(){
	// 	fmt.Println("I'm mr.anonymous")
	// }()

	// ex9 func expression
	// y:=func ()  {
	// 	for i:=0;i<10;i++{
	// 		fmt.Println(i)
	// 	}
	// }
	// y()

	// ex9 func return
	// y:=bar()
	// fmt.Println(y())

	// ex10 callback
	// result := printSquare(Square, 5)
	// fmt.Println(result)

	// ex11 closure
	// x:=powr(2)
	// fmt.Println(x())
	// fmt.Println(x())
	// fmt.Println(x())
	// fmt.Println(x())
	// fmt.Println(x())
	// fmt.Println(x())
	// fmt.Println(x())
	// fmt.Println(x())

	// ex12 wrapper
	// timeFunc(doWork)

}

// func doWork() {
// 	for i := 0; i < 100; i++ {
// 		fmt.Println(i)
// 	}
// }
// func timeFunc(f func()) {
// 	start := time.Now()
// 	f()
// 	elapsed := time.Since(start)
// 	fmt.Println(elapsed)
// }
