package base62

import (
    "errors"
    "math"
    "strings"
)

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

func Decode(encoded string) (uint64, error){
    var x uint64

    for i, symbol := range encoded {
        alphabeticPosition := strings.IndexRune(alphabet,  symbol )

        if alphabeticPosition == -1 {
            return uint64(alphabeticPosition), errors.New("Invalid character" + string(symbol))
        }

        x += uint64(alphabeticPosition) * uint64(math.Pow(float64(length), float64(i)))
    }

    return x, nil
}
