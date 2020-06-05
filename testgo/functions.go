package main

import (
    "fmt"
)

func add(x int,y int) int {
    return x + y
}

func add2(x,y int) int {
    return x+y
}

func swap(x,y string) (string,string) {
    return y,x
}

func test(x,y int) (s int) {
    s = x * y
    return
}

func main(){
    fmt.Println(add(2,4))
    fmt.Println(add2(2,4))

    x,y := swap("a","b")
    fmt.Println(x,y)

    fmt.Println(test(1,2))
}
