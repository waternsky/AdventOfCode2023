package main

import (
	"fmt"
	"os"
    "strconv"
)

func main() {
    
    paresedData, _ := os.ReadFile("./input.1.1.txt")

    first, last := 0, 0
    var sum int64 = 0
    dFound := false
    for _, ch := range paresedData {
        chs := string(ch)
        i, err := strconv.ParseInt(chs, 10, 64) 

        if chs == "\n" {
            sum += int64(last + 10 * first)
            first = 0
            last = 0
            dFound = false
        }

        if err != nil {
            continue
        } 

        if dFound {
            last = int(i)
        } else {
            first = int(i)
            last = int(i)
            dFound = true
        }
    }

    fmt.Println(sum)

}
