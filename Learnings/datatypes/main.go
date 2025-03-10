package main

import "fmt"

func main() {
    var a int = 10
    fmt.Println("Integer (int) value is set to:", a)

    var b float64 = 3.14
    fmt.Println("Floating-point (float64) value is set to:", b)

    var c string = "Hello, Go!"
    fmt.Println("String value is set to:", c)

    var d bool = true
    fmt.Println("Boolean value is set to:", d)

    var e complex128 = 1 + 2i
    fmt.Println("Complex number value is set to:", e)

    var f byte = 'A'
    fmt.Println("Byte (as character) value is set to:", string(f))

    var g rune = 'G'
    fmt.Println("Rune (as character) value is set to:", string(g))

    var h int8 = -128
    fmt.Println("Int8 value is set to:", h)

    var i uint = 42
    fmt.Println("Unsigned Int (uint) value is set to:", i)

    var j uint16 = 65535
    fmt.Println("Unsigned Int16 value is set to:", j)

    var k uint32 = 4294967295
    fmt.Println("Unsigned Int32 value is set to:", k)

    var l uint64 = 18446744073709551615
    fmt.Println("Unsigned Int64 value is set to:", l)
}
