package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// json marshal
type Color struct {
	Id     int
	Name   string
	Colors []string
}

// json unmarshal
type Animal struct {
	Name  string
	Order string
}

func main() {
	//json marshal
	g := Color{
		Id:     1,
		Name:   "jkay",
		Colors: []string{"green", "yellow"},
	}
	b, err := json.Marshal(g)
	if err != nil {
		fmt.Println("error:", err)
	}
	os.Stdout.Write(b)
	fmt.Println("\n",string(b))

	//json unmarshal
	var jsonBlob = []byte(`[
		{"Name":"Play","Order":"blah"}
	]`)
	var animal []Animal
	err = json.Unmarshal(jsonBlob, &animal)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("\n%+v", animal)
}
