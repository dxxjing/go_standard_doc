package main

import "fmt"

type area struct {
    length int
    height int
}

var m map[string]area

func main(){
    m1 := make(map[string]area)
    m1["a"] = area{1,2}
    fmt.Println(m1)
    /* 
    m2 := make(map[string]area)
    m2 = m{"a":area{1,2},"b":area{3,4}}
    fmt.Println(m2)
    */

    m3 := map[string]area{"a":area{1,2},"b":area{3,4}}
    fmt.Println(m3)
    
    m4 := make(map[string]area)
    m4 = map[string]area{"a":area{1,2},"b":area{3,4}}
    fmt.Println(m4)
    
    m5 := map[string]area{"a":{1,2},"b":{3,4}}
    fmt.Println(m5)

    m5["a"] = area{11,22}
    fmt.Println(m5)

    if _,ok := m5["b"];ok {
        fmt.Println("existed")
    }
}
