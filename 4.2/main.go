package main

import (
	"fmt"
	"os"
	"strings"
    "regexp"
    "strconv"
)

func main() {

    parseData, _ := os.ReadFile("../4.1/input.4.1.txt")

    parseAsStr := strings.TrimSpace(string(parseData))
    /*
    parseAsStr = strings.TrimSpace(`
    Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
    Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
    Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
    Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
    Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
    Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11
    `)
    */
    splitData := strings.Split(parseAsStr, "\n")

    m := make(map[int]int)

    for i := 1; i <= len(splitData); i++ {
        m[i] = 1
    }

    fmt.Println("Card: ", splitData[0], "has", winPerCard(splitData[0]))
    for idx := range splitData {
        cardNum := idx + 1
        for j := 1; j <= int(winPerCard(splitData[idx])); j++ {
            if cardNum + j <= len(splitData) { 
                //fmt.Println("Cardnum: ", cardNum+j, " : ", m)
                m[cardNum+j] += m[cardNum]
            }
        }        
    }
    fmt.Println(m)

    var sum int = 0
    for _, j := range m {
        sum += j
    }

    fmt.Println(sum)

}

func winPerCard(card string) int64 { 
    initial := strings.Split(card, ":")[1]
    winners := strings.Split(initial, "|")[0]
    drawn := strings.Split(initial, "|")[1]

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
    fmt.Println("Winner: ", winNumarr)
    fmt.Println("Drawn: ", drawnNumarr)

    return int64((numOfMatches(winNumarr, drawnNumarr)))
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
