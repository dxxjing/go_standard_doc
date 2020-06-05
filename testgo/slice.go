package main

import "fmt"

func main(){
    p := []int{1,2,3,4}
    fmt.Println(p)

    for i := 0; i < len(p); i++ {
        fmt.Println(p[i])
    }

    p2 := p[1:3]
    fmt.Println(p2)
    p3 := p[:3]
    fmt.Println(p3)


    a := make([]int,5)
    printS("a",a)

    b := make([]int,0,5)
    printS("b",b)

    c := a[1:]
    printS("c",c)


    var d []int
    if d == nil {
        printS("d",d)
    }

    sl := make([]int,5)
    sl = append(sl,1)
    printS("sl",sl)

    sl = append(sl,2,3,4,5,6,7)
    printS("sl",sl)

    sa := make([]int,0)
    sa = append(sa,1,2,3)
    printS("sa",sa)

}

func printS(s string,a []int){
    fmt.Printf("%s len=%d,cap=%d %v\n",s,len(a),cap(a),a)
}
