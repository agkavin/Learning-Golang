package main

import (
    "fmt"
    "Golang_Learnings/pkg/calc" // Import calc
)

func main() {
    // Using public function Sum from calc
    sumResult := calc.Sum(10, 20)
    fmt.Println("Sum:", sumResult)

    // Accessing public variable Version from calc
    fmt.Println("Version:", calc.Version)

    // Private variables/functions like logMessage or sum cannot be accessed outside the package
}
