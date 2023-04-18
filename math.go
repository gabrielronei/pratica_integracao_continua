package main

import "fmt"

func main() {
     fmt.Println(Sum(10, 40))
     fmt.Println(Subtract(10, 40))
}

func Sum(a int, b int) int {

     return a + b
}

func Subtract(a int, b int) int {

     return a - b
}