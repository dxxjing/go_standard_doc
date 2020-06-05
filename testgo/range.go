package main

import "fmt"

func main(){
    s := []int{1,2,3,4,5,6,7,8,9}

    for k,v := range s {
        fmt.Printf("%d=>%d\n",k,v)
    }
    
    for v := range s {
        fmt.Printf("%d\n",v)
    }
    
    for _,v := range s {
        fmt.Printf("%d\n",v)
    }
    
    for k,_ := range s {
        fmt.Printf("%d\n",k)
    }
}
