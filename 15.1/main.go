package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
    parseData, _ := os.ReadFile("./input.15.1.txt")
    parseAsStr := string(parseData)

    splitData := strings.Split(strings.TrimSpace(parseAsStr), ",")

    ans := 0
    for i := range splitData {
        ans += algo(splitData[i])
    }
    fmt.Println(ans)
}

func algo(s string) int {
    val := 0
    for i := range s {
        ascii := int(s[i]) // rune 
        val += ascii
        val *= 17
        val = val % 256
    }
    return val
}
