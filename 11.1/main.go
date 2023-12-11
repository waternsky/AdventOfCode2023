package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strings"
)

func main() {
    parseData, _ := os.ReadFile("./input.11.1.txt")
    parseAsStr := string(parseData)
    /*
    parseAsStr = `
...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....
    `
    */
    splitData := strings.Split(strings.TrimSpace(parseAsStr), "\n")
    //fmt.Println(splitData)

    re := regexp.MustCompile("#")

    expandRow := make([]string, 0, 2 * len(splitData))
    for i := 0; i < len(splitData); i++ {
        if re.FindString(splitData[i]) == "" {
            expandRow = append(expandRow, splitData[i])
        }
        expandRow = append(expandRow, splitData[i])
    }

    expandCol := make([]string, 0, 2 * len(splitData[0]))
    tmp := make([]int, 0, 2 * len(splitData[0]))
    for i := 0; i < len(splitData); i++ {
        if re.FindAllStringIndex(splitData[i], -1) == nil {
            continue
        }
        elm := re.FindAllStringIndex(splitData[i], -1)
        in:
        for j := 0; j < len(elm); j++ {
            if contains(elm[j][0], tmp) {
                continue in
            }
            tmp = append(tmp, elm[j][0])
        }
    }

    tmp2 := make([]int, 0, 2 * len(splitData[0]))
    for i := 0; i < len(splitData[0]); i++ {
        if contains(i, tmp) {
            continue
        }   
        tmp2 = append(tmp2, i)
    }

    for i := 0; i < len(expandRow); i++ {
        expandCol = append(expandCol, insert(expandRow[i], tmp2))
    }

    //fmt.Println(expandCol)

    m := make([][2]int, 0, len(expandCol) * len(expandCol[0]))
    for i := 0; i < len(expandCol); i++ {
        elm := re.FindAllStringIndex(expandCol[i], -1)
        if elm == nil {
            continue 
        }
        for j := 0; j < len(elm); j++ {
            m = append(m, [2]int{i, elm[j][0]})
        }
    }
    //fmt.Println(m)
   
    ans := 0
    for i := 0; i < len(m); i++ {
        for j := i+1; j < len(m); j++ {
            //fmt.Println("Distance between ", i+1, "and ", j+1, " is ", manhattan(m[i], m[j]))
            ans += manhattan(m[i], m[j])
        } 
    }

    fmt.Println(ans)
}

func manhattan(x [2]int, y [2]int) int {
    return int(math.Abs(float64(y[1] - x[1])) + math.Abs(float64(y[0] - x[0])))
}

func insert(s string, arr []int) string {
    ans := ""
    for i := 0; i < len(s); i++ {
        if contains(i, arr) {
            ans = ans + "."
        }
        ans = ans + string(s[i])
    }
    return ans
}

func contains(elm int, arr []int) bool {
    for i := range arr {
        if arr[i] == elm {
            return true
        }
    }
    return false
}
