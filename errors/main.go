//ex1

// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"log"
// )

// type person struct {
// 	First   string
// 	Last    string
// 	Sayings []string
// }

// func main() {
// 	p1 := person{
// 		First:   "James",
// 		Last:    "Bond",
// 		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
// 	}

// 	bs, err := json.Marshal(p1)
// 	if err!= nil{
// 		// fmt.Println("error:",err)
// 		log.Fatalln("error:",err)
// 	}
// 	fmt.Println(string(bs))

// }

//ex2

// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"log"
// )

// type person struct {
// 	First   string
// 	Last    string
// 	Sayings []string
// }

// func main() {
// 	p1 := person{
// 		First:   "James",
// 		Last:    "Bond",
// 		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
// 	}

// 	bs, err := toJSON(p1)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println(string(bs))

// }

// // toJSON needs to return an error also
// func toJSON(a interface{}) ([]byte, error) {
// 	bs, err := json.Marshal(a)
// 	if err != nil {
// 		return []byte{}, fmt.Errorf("There was an error in toJSON: %v", err)
// 	}
// 	return bs, nil
// }

//ex3

// package main

// import (
// 	"fmt"
// )

// type customErr struct {
// 	info string
// }

// func (ce customErr) Error() string {
// 	return fmt.Sprintf("haha here is the error:%v",ce.info)
// }

// func main() {
// 	c1 := customErr{
// 		info: "error",
// 	}
// 	foo(c1)
// }

// func foo(e error) {
// 	fmt.Println("lalaal",e)
// }

// ex4

package main

import (
	"errors"
	"fmt"
	"log"
)

type sqrtError struct {
	lat  string
	long string
	err  error
}

func (se sqrtError) Error() string {
	return fmt.Sprintf("math error: %v %v %v", se.lat, se.long, se.err)
}

func main() {
	_, err := sqrt(-10.23)
	if err != nil {
		log.Println(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		e:=errors.New("need more error")
		return 0,sqrtError{
			"50.2289 N",
            "99.4656 W",
			e,
		}
	}
	return 42, nil
}
