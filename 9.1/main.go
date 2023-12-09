package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
    parseData, _ := os.ReadFile("./input.9.1.txt")
    parseAsStr := string(parseData)

    splitData := strings.Split(strings.TrimSpace(parseAsStr), "\n")

    var sum int64 = 0
    for i := range splitData {
        sum += getNextNum(splitData[i])
    }

    fmt.Println(sum)
}

func getNextNum(inp string) int64 {
    arr := strings.Split(strings.TrimSpace(inp), " ")
    intArr := make([]int64, 0, len(arr))
    for i := range arr {
        integer, _ := strconv.ParseInt(arr[i], 10, 64)
        intArr = append(intArr, integer)
    }

    x := make([][]int64, 0, len(intArr))
    start := intArr
    x = append(x, start)
    for !allZeros(start) {
        start = getDiff(start)
        x = append(x, start)
    }
    
    var num int64 = 0
    for i := 0; i < len(x); i++ {
        tmp := x[len(x) - i - 1]
        num += tmp[len(tmp) - 1]
    }

    return num
}

func getDiff(arr []int64) []int64 {
    ans := make([]int64, 0, len(arr) - 1)
    for i := 0; i < len(arr) - 1; i++ {
        ans = append(ans, arr[i+1] - arr[i])
    }
    return ans
}

func allZeros(arr []int64) bool {
    ans := true
    for i := range arr {
        ans = ans && (arr[i] == 0)
    }
    return ans
}
