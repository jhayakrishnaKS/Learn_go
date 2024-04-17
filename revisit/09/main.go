package main

import (
	main1 `common/revisit/08`
	"fmt"
)

type hello struct {
    main1.Gender 
    age int
}

func main() {
    ins := hello{
        Gender: main1.Gender{Male: false},
        age:    10,
    }

	fmt.Printf("%+v",ins)
}
