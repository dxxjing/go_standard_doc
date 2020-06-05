package main

import (
	"fmt"
	"os"
    "io"
)

func ReadFile(filename string) (string,error) {
    f,err := os.Open(filename)
    if err != nil {
        return "",err
    }
    defer f.Close()

    var result []byte
    buf := make([]byte,10)
    for {
        n,err := f.Read(buf[0:])
        result = append(result,buf[:n]...)
        if err != nil {
            if err == io.EOF {
                break;
            }
            return "",err
        }
    }
    return string(result),nil
}

func main(){

   s,err :=  ReadFile("test.txt")
   if err != nil {
        fmt.Println(err)
   }
   fmt.Printf("%s\n",s)

}
