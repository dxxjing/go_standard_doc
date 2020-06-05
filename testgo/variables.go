package main

import (
    "fmt"
)

var x,y bool
var z string
func main(){

    var a int
    fmt.Println(x,y,z,a)

    var i,j = 3,"test"
    fmt.Println(i,j)


    t,x := 3,true
    fmt.Println(t,x)
}
