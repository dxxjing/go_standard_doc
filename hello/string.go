package main

import "fmt"

func Compare(a,b []byte) int {
    for i := 0; i < len(a) && i < len(b); i++ {
        if a[i] > b[i] {
            return 1
        }else if a[i] < b[i] {
            return -1
        }
    }

    switch {
        case len(a) > len(b) : return 1
        case len(a) < len(b) : return -1
    }
    return 0
}

func main(){
    a := []byte("abc")
    b := []byte("acd")
    fmt.Println(Compare(a,b))

}
