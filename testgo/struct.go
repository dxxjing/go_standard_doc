package main

import "fmt"

type s struct {
    x int
    y string
}

func main(){
    s1 := s{1,"test"}
    fmt.Println(s1)

    s1.x = 100
    fmt.Println(s1)

    p := &s{199,"hello"}
    p.x = 99
    fmt.Println(p,*p)


    //v1 := s{1}  error 

    v2 := s{x:1,y:"dd"}
    v3 := s{y:"hh"}
    fmt.Println(v2,v3)
}
