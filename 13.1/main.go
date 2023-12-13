package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
    
    parseData, _ := os.ReadFile("./input.13.1.txt")
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
    fmt.Println(transpose(splitData[0]))

    total := 0
    for i := range splitData {
        fmt.Println(findMirror(splitData[i]), findMirror(transpose(splitData[i])))
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

func findMirror(valley string) int {
    arr := strings.Split(valley, "\n")
    for i := 0; i < len(arr) - 1; i++ {
        ans := true
        for j := 0; j <= i && j < len(arr) - i - 1; j++ {
            ans = ans && arr[i-j] == arr[i+1+j]
        }
        if ans {
            return i+1
        }
    }
    return 0 
}
