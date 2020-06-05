package main

import "fmt"

func main(){
    sum := 0
    for i := 0;i < 10; i++ {
        sum += i
    }
    fmt.Println(sum)



    s := 1
    for ;s < 10; {
        s += s
    }
    fmt.Println(s)


    //like while
    x := 1
    for x < 5 {
        x += x
    }
    fmt.Println(x)

    /*
    //死循环
    for {
    
    }
    */
}
