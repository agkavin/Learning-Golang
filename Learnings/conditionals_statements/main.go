package main

import (
    "fmt"
    "Golang_Learnings/pkg/conditionals"
)

func package_class() {
    fmt.Println("Which conditionals do you want to see?")
    fmt.Println("1. if types")
    fmt.Println("2. switch")
    var choice int
    fmt.Scan(&choice)
    if choice == 1 {
        conditionals.Test_run()
    } else if choice == 2 {
        conditionals.Switch_example()
    }
}

func cond() {
    if 3 == 7 {
        fmt.Println("its 7, Thala for a reason 🚁")
    }else {
        fmt.Println("3 !! Aint 7 😠")
    }
}

func main() {
    fmt.Println("Conditionals in Go, as a Package")
    package_class()
    fmt.Println("_____")
    fmt.Println("Function defined in this file")
    cond()
}