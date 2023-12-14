package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
    parseData, _ := os.ReadFile("./input.14.1.txt")
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
    splitData := strings.Split(strings.TrimSpace(parseAsStr), "\n")

    ans := 0
    for i := 0; i < len(splitData[0]); i++ {
        ballIdx := i
        ballCount := 0
        in:
        for j := range splitData {
            if string(splitData[j][ballIdx]) == "#" {
                break in
            }
            if string(splitData[j][ballIdx]) == "O" {
                ballCount++
            }
        }
        for k := 0; k < ballCount; k++ {
            ans += len(splitData) - k
        }
    }

    re := regexp.MustCompile("#")
    for j := 0; j < len(splitData) - 1; j++ {
        for _, v := range re.FindAllStringIndex(splitData[j], -1) {
            ballCount := 0
            in2:
            for i := j+1; i < len(splitData); i++ {
                if string(splitData[i][v[0]]) == "#" {
                    break in2
                }
                if string(splitData[i][v[0]]) == "O" {
                    ballCount++
                }
            }
            for k := 0; k < ballCount; k++ {
                ans += len(splitData) - j - k - 1
            }
        }
    }


    fmt.Println(ans)
}
