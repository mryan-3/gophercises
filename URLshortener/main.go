package main

import (
    "fmt"
   "strings"
)

func main(){

    var thre = Encode(1029497)
    fmt.Println(thre)


}

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const length = uint64(len(alphabet))

func Encode(x uint64) string{
    var encodeBuilder strings.Builder
    encodeBuilder.Grow(11)

    for ; x > 0; x = x / length {
        encodeBuilder.WriteByte(alphabet[(x % length)])
    }
    return encodeBuilder.String()
}
