package main

import "testing"

func TestSum(t *testing.T){
	X:=mySum(2,3)
	if X!=5{
		t.Error("Expected",5,"Got",X)
	}
}