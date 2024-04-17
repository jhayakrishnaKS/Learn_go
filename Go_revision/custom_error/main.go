package main

import (
    "fmt"
)

type CustomError struct {
    Code    int
    Message string
}

func (e *CustomError) Error() string {
    return fmt.Sprintf("Error %d: %s", e.Code, e.Message)
}

func someFunction() error {
    return &CustomError{Code: 500, Message: "Something went wrong cause you have nobrain dumb"}
}

func main() {
    err := someFunction()
    if err != nil {
        fmt.Println("Error:", err)
    }
}
