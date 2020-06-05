package main

import (
    "fmt"
    "strings"
)


func main(){

    r := strings.NewReader("hello")
    s := make([]byte,8)
    
    n,err := r.Read(s)
    fmt.Println(n, err, s)
}
