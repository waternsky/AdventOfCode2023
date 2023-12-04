package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {

    parseData, _ := os.ReadFile("./input.4.1.txt")

    parseAsStr := strings.TrimSpace(string(parseData))

    splitData := strings.Split(parseAsStr, "\n")

    var ans int64 = 0
    fmt.Println(splitData[len(splitData)-2])
    for idx := range splitData {
        ans += winPerCard(splitData[idx])
    }

    fmt.Println(ans)
}

func winPerCard(card string) int64 { 
    initial := strings.Split(card, ":")[1]
    winners := strings.Split(initial, "|")[0]
    drawn := strings.Split(initial, "|")[1]
    fmt.Println("winner: ", winners)
    fmt.Println("drawn: ", drawn)

    arrW := strings.Join(strings.Split(strings.TrimSpace(winners), " "), " ")

    re := regexp.MustCompile("[ ]+")

    ans := re.Split(arrW, -1)

    winNumarr := make([]int64, 0, len(ans))
    for i := range ans {
        in, _ := strconv.ParseInt(ans[i], 10, 64)
        winNumarr = append(winNumarr, in)
    }
    
    arrD := strings.Join(strings.Split(strings.TrimSpace(drawn), " "), " ")

    ans2 := re.Split(arrD, -1)

    drawnNumarr := make([]int64, 0, len(ans2))
    for i := range ans2 {
        in, _ := strconv.ParseInt(ans2[i], 10, 64)
        drawnNumarr = append(drawnNumarr, in)
    }

    fmt.Println(winNumarr)
    fmt.Println(drawnNumarr)
    fmt.Println(numOfMatches(winNumarr, drawnNumarr))

    return int64(math.Pow(2, float64(numOfMatches(winNumarr, drawnNumarr) - 1)))
}

func numOfMatches(a []int64, b []int64) int {
    ans := 0
    for i := 0; i < len(a); i++ {
        for j := 0; j < len(b); j++ {
            if a[i] == b[j] {
                ans++
            }
        }
    }
    return ans
}


