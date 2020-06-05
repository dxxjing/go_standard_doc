package main

import "fmt"

func main(){

    var a [2]int
    a[0] = 1
    a[1] = 2
    fmt.Println(a)
    /* error 
    a = [1,2]
    fmt.Println(a)
    */
}
