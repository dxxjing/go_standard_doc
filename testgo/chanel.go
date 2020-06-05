package main

import (
    "fmt"
    "time"
)

func sum(s []int,c chan int){
    sum := 0
    for _,v := range s {
        sum += v
    }
    c <- sum
}

func add(c chan int){
    defer close(c)
    for i := 0; i < 10;i++ {
        c <- i
    }
}

func over(c chan int){

    select {
        case v := <- c :
            fmt.Println(v)
            return
        case <-time.After(3*time.Second):
            fmt.Println("timeout")
    }
}

func main(){
    a := []int{1,2,3,4,5,6,7,8,9}
    c := make(chan int)
    go sum(a[:len(a)/2],c)
    go sum(a[len(a)/2:],c)
    x,y := <-c,<-c
    fmt.Println(x+y)
    /*
    c2 := make(chan int,2)
    c2 <- 3
    c2 <- 4
    c2 <- 5
    fmt.Println(<-c2)
    */

    c4 := make(chan int,20)
    go add(c4)
    for i := range c4 {
        fmt.Println(i)
    }


}
