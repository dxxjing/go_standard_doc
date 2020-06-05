package main

import (
	"fmt"
	"net/http"
)


func main(){
    fmt.Println(http.StatusOK)
    fmt.Println(http.DefaultMaxHeaderBytes)

    pErr := &http.ProtocolError{"test"}
    fmt.Println(pErr.Error())

    fmt.Println(http.CanonicalHeaderKey("accept-encoding"))

    fmt.Println(http.DetectContentType([]byte("mp3")))
}
