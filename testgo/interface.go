package main

import "fmt"

type Adder interface{
    Add() float64
}


type Flt float64

func (f Flt) Add() float64 {
    return float64(f * 0.1)
}

type Vertex struct{
    x float64
    y float64
}

func (v *Vertex) Add() float64 {
    return v.x +v.y
}


func main(){

    var i Adder
    f := Flt(0.1)
    v1 := &Vertex{0.1,0.2}
    //v2 := Vertex{0.1,0.2}  不能将v2赋值给i 并调用Add()


    i = f
    fmt.Println(i.Add())

    i = v1
    fmt.Println(i.Add())

}


