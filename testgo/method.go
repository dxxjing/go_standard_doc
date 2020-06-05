package main

import "fmt"

type Vertex struct{
    x int
    y int
}

func (v *Vertex) add() int {
    return v.x + v.y
}

func (v *Vertex) swap(){
    v.x,v.y = v.y,v.x
}

func main(){

    v1 := &Vertex{1,2}
    fmt.Println(v1.add())

    v2 := Vertex{1,3}
    fmt.Println(v2.add())


    v3 := &Vertex{1,2}
    v3.swap()
    fmt.Println(v3)

    v4 := Vertex{1,2}
    v4.swap()
    fmt.Println(v4)

}
