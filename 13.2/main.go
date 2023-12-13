package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
    
    parseData, _ := os.ReadFile("../13.1/input.13.1.txt")
    parseAsStr := string(parseData)
    /*
    parseAsStr = `
#.##..##.
..#.##.#.
##......#
##......#
..#.##.#.
..##..##.
#.#.##.#.

#...##..#
#....#..#
..##..###
#####.##.
#####.##.
..##..###
#....#..#
    `
    parseAsStr = `
..##.###..###.##.
##..#..#..#..#..#
###..#..##..#.###
##..##########..#
###.####..####.##
..#...#.##.#...#.
##..#.#....#.#..#
..#.##.####.##.#.
..##.########.##.
    `
    */
    splitData := strings.Split(strings.TrimSpace(parseAsStr), "\n\n")

    total := 0
    for i := range splitData {
        //fmt.Println(findMirror(splitData[i]), findMirror(transpose(splitData[i])))
        total += findMirror(splitData[i]) * 100 + findMirror(transpose(splitData[i]))
    }
    fmt.Println(total)
}

func transpose(valley string) string {
    arr := strings.Split(valley, "\n")
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

func diffCount(s string, t string) int {
    if len(s) != len(t) {
        return -1000000
    }
    count := 0
    for i := 0; i < len(s); i++ {
        if string(s[i]) != string(t[i]) {
            count++
        }
    }
    return count
}

func findMirror(valley string) int {
    arr := strings.Split(valley, "\n")
    for i := 0; i < len(arr) - 1; i++ {
        ans := true
        dcount := 0
        for j := 0; j <= i && j < len(arr) - i - 1; j++ {
            ans = ans && arr[i-j] == arr[i+1+j]
            dcount += diffCount(arr[i-j], arr[i+1+j])
        }
        // fmt.Println("for i = ", i, " dcount = ", dcount)
        if dcount == 1 {
            return i+1
        }
    }
    return 0 
}
