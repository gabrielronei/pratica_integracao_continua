package main

import "fmt"

func main() {
     fmt.Println(Sum(10, 40))
     fmt.Println(Subtract(10, 40))
}

func sum(a int, b int) int {

     return a + b
}

func subtract(a int, b int) int {

     return a - b
}