package conditionals

import (
    "fmt"
    "math/rand"
)

func Switch_example() {
    i := rand.Int31n(10)

    switch i {
    case 0:
        fmt.Print("zero...")
    case 1:
        fmt.Print("one...")
    case 2:
        fmt.Print("two...")

    default:
        fmt.Print("ran...")

    }

    fmt.Println("ok")
}