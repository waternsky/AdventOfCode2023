package main

import (
	"fmt"
	"os"
	"strings"
    "regexp"
)

func main() {
    parseData, _ := os.ReadFile("../14.1/input.14.1.txt")
    parseAsStr := string(parseData)
    /*
    parseAsStr = `
O....#....
O.OO#....#
.....##...
OO.#O....O
.O.....O#.
O.#..O.#.#
..O..#O..O
.......O..
#....###..
#OO..#....
    `
    */
    parseAsStr = strings.TrimSpace(parseAsStr)

    tmp := -1
    cache := make([]string, 0, 100)
    x := 0
    for i := 0; i < 1000000000; i++ {
        parseAsStr = cycle(parseAsStr)
        x++
        if isIn(parseAsStr, cache) >= 0 {
            tmp = isIn(parseAsStr, cache)
            break
        }
        cache = append(cache, parseAsStr)
    }
    //fmt.Println(tmp, x) // 123, 150 cycle length 16, 150 == 124 cycle and so on
    
    //cycleStart := tmp + 1
    cycleLength := len(cache) - tmp
    fmt.Println(calculateLoad(cache[tmp+(1000000000-tmp-1) % cycleLength]))
}

func calculateLoad(parseAsStr string) int {
    x := strings.Split(parseAsStr, "\n")
    ans := 0
    for i := range x {
        re := regexp.MustCompile("O")
        tmp := re.FindAllStringIndex(x[i], -1)
        ans += len(tmp) * (len(x) - i)
    }

    return ans
}

func isIn(check string, arr []string) int {
    for i := range arr {
        if check == arr[i] {
            return i
        }
    }
    return -1
}

func cycle(s string) string {
    x := pullUp(s)
    x = pullUp(rotateClockWise(x))
    x = pullUp(rotateClockWise(x))
    x = pullUp(rotateClockWise(x))
    return rotateClockWise(x)
}

func pullUp(s string) string {
    x := strings.Split(s, "\n")
    a := make([]string, 0, len(x))
    for j := 0; j < len(x[0]); j++ {
        sqIdx := make([]int, 0, len(x))
        col := ""
        for i := range x {
            if string(x[i][j]) == "#" {
                sqIdx = append(sqIdx, i)
            }
            col += string(x[i][j])
        }
        ballCol := make([]int, len(sqIdx)+1)
        idx := 0
        for i := range x {
            if string(x[i][j]) == "O" {
                ballCol[idx]++
            }
            if string(x[i][j]) == "#" {
                idx++
                continue
            }
        }
        //fmt.Println(ballCol)
        //fmt.Println(sqIdx)
        //fmt.Println(col)
        oCount := 0
        sqCount := 0
        ans := ""
        for i := 0; i < len(x); i++ {
            if len(sqIdx) > 0 && i < sqIdx[0] && oCount < ballCol[0] {
                ans += "O"
                oCount++
            } else if len(sqIdx) > 0 && i > sqIdx[0] && sqCount < len(sqIdx) && oCount + sqIdx[sqCount] < ballCol[sqCount] {
                oCount++
                ans += "O"
            } else if len(sqIdx) > 0 && i > sqIdx[0] && oCount < ballCol[sqCount] {
                oCount++
                ans += "O"
            } else if len(sqIdx) > 0 && sqCount < len(sqIdx) && i == sqIdx[sqCount] {
                ans += "#"
                sqCount++
                oCount = 0
            } else if len(sqIdx) > 0 && sqCount < len(sqIdx) && i > sqIdx[sqCount] && oCount + sqIdx[sqCount] < ballCol[sqCount + 1] {
                ans += "O"
                oCount++
            } else if len(sqIdx) == 0 && oCount < ballCol[sqCount] {
                ans += "O"
                oCount++
            } else {
                ans += "."
            }
        }
        // fmt.Println(ans)
        a = append(a, ans)
    }

    return transpose(strings.Join(a, "\n"))
}

func transpose(s string) string {
    arr := strings.Split(s, "\n")
    ans := make([]string, 0, len(arr[0]))
    for i := 0; i < len(arr[0]); i++ {
        col := ""
        for j := range arr {
            col += string(arr[j][i])
        }
        ans = append(ans, col)
    }

    return strings.Join(ans, "\n")
}

func reverse(s string) string {
    ans := ""
    for i := range s {
        ans += string(s[len(s)-1-i])
    }
    return ans
}

func rotateClockWise(s string) string {
    arr := strings.Split(s, "\n")
    ans := make([]string, 0, len(arr[0]))
    for i := 0; i < len(arr[0]); i++ {
        col := ""
        for j := range arr {
            col += string(arr[j][i])
        }
        ans = append(ans, reverse(col))
    }

    return strings.Join(ans, "\n")
}
