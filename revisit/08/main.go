package main1

import (
	
	"fmt"
)

type Gender struct {
	Male bool
}

type Element struct {
	Name   string
	Age    int
	Gender Gender
}

func main() {
	elements := []Element{
		{
			Name: "jsj",
			Age:  20,
			Gender: Gender{
				Male: true,
			},
		},
	}

	for i, v := range elements {
		fmt.Printf("%d: %+v\n", i, v)
	}
	
	
}
