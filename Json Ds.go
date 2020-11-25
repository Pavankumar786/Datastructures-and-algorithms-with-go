package main

import (
	"fmt"
	"encoding/json"
)

type Message struct {
    Name string
    Body string
    Time int64
}

func main() {
	m := Message{"Interface", "Empty Interface", 1556747623}
	b, err := json.Marshal(m)
	fmt.Printf("err = %v\n", err)
	fmt.Printf("b = %T%+v\n", b, b)
	fmt.Printf("b = %T%s\n", b, b)
}


output:
err = <nil>
b = []uint8[123 34 78 97 109 101 34 58 34 73 110 116 101 114 102 97 99 101 34 44 34 66 111 100 121 34 58 34 69 109 112 116 121 32 73 110 116 101 114 102 97 99 101 34 44 34 84 105 109 101 34 58 49 53 53 54 55 52 55 54 50 51 125]
b = []uint8{"Name":"Interface","Body":"Empty Interface","Time":1556747623}