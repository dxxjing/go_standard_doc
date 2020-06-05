package main 

import "fmt"

func main(){
    x := true
    if x {
        fmt.Println("true")
    }else{
        fmt.Println("false")
    }


    if y := 1; y < 3 {
        fmt.Println("x lt 3")
    }
    //fmt.Println(y)  undefine 'y'
}
