package main

import "fmt"
import "runtime"

func main(){
    
    switch os := runtime.GOOS; os {
        case "darwin" :
            fmt.Println("OS .x")
        case "linux":
            fmt.Println("unix")
        default:
            fmt.Println("unknown")
    }
}
