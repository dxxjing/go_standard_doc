package main

import "fmt"

func a(){
    fmt.Println("in a")
}

func b(s string){
    defer trace(s)
    fmt.Println("in b")
}

func trace(s string){
    fmt.Println("trace:",s)
}

func main(){
    defer b("test")
    a()
}
