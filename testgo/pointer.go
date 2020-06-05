package main

import "fmt"

func main(){
    
    i := 2
    var p *int

    p = &i
    fmt.Println(*p)

    *p = 3
    fmt.Println(i)


}
