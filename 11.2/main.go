package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strings"
)

func main() {
    parseData, _ := os.ReadFile("../11.1/input.11.1.txt")
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

    tp := make([]int, 0, len(splitData))
    for i := 0; i < len(splitData); i++ {
        if re.FindString(splitData[i]) == "" {
            tp = append(tp, i)
        }
    }

    //fmt.Println(tp)

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

    //fmt.Println(tmp2)

    m := make([][2]int, 0, len(splitData) * len(splitData[0]))
    for i := 0; i < len(splitData); i++ {
        elm := re.FindAllStringIndex(splitData[i], -1)
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
            ans += manhattan(m[i], m[j], tp, tmp2, 1000000)
        } 
    }

    fmt.Println(ans)
}

func manhattan(x [2]int, y [2]int, xspace []int, yspace []int, xpf int) int {
    xfactor, yfactor := 0, 0
    for i := 0; i < len(xspace); i++ {
        if x[0] < y[0] && x[0] < xspace[i] && y[0] > xspace[i] {
            xfactor++
        }
        if x[0] > y[0] && y[0] < xspace[i] && x[0] > xspace[i] {
            xfactor++
        }
    }
    for i := 0; i < len(yspace); i++ {
        if x[1] < y[1] && x[1] < yspace[i] && y[1] > yspace[i] {
            yfactor++
        }
        if x[1] > y[1] && y[1] < yspace[i] && x[1] > yspace[i] {
            yfactor++
        }
    }
    return int(math.Abs(float64(y[1] - x[1])) + math.Abs(float64(y[0] - x[0]))) + xfactor * (xpf - 1) + yfactor * (xpf - 1)
}

func insert(s string, arr []int, ntime int) string {
    ans := ""
    for i := 0; i < len(s); i++ {
        if contains(i, arr) {
            for j := 1; j < ntime; j++ {
                ans = ans + "."
            }
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
