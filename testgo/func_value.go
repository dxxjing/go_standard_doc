package main

import "fmt"

func main(){
    f := func (x,y int)int {
        return x+y
    }
    fmt.Println(f(1,2))


    f1 := adder()
    for i := 1;i < 5; i++ {
        fmt.Println(f1(i))
    }
}


func adder() func(int) int {
    sum := 0
    return func(x int) int {
        sum += x
        return sum
    }
}

