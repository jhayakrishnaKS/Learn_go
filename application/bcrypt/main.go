package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	s := `password`
	bs, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(s)
	fmt.Println(string(bs))

	login:=`password`
	err=bcrypt.CompareHashAndPassword(bs, []byte(login))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("you logged in")
}
