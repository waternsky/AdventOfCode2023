package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
    parseData, _ := os.ReadFile("./input.21.1.txt")
    parseAsStr := string(parseData)
    /*
    parseAsStr = `
...........
.....###.#.
.###.##..#.
..#.#...#..
....#.#....
.##..S####.
.##..#...#.
.......##..
.##.#.####.
.##..##.##.
...........
    `
    */
    parseAsStr = string(parseAsStr)
    splitData := strings.Split(parseAsStr, "\n")

    var sIndex [2]int
    for i := range splitData {
        chk := regexp.MustCompile("S").FindStringIndex(splitData[i])
        if chk != nil {
            sIndex = [2]int{i, chk[0]} 
            break
        }
    }

    count := 0
    start := [][2]int{sIndex}
    for count < 64 {
        //fmt.Println(start)
        start = nbrwalk(start, splitData) 
        count++
    }

    fmt.Println(len(start))
}

func nbrwalk(node [][2]int, inp []string) [][2]int {
    
    var visited [][2]int
    for i := range node {
        visited = merge(nbrs(node[i], inp), visited) 
    }
    //fmt.Println("Visited", visited)
    return visited
}

func merge(x [][2]int, y [][2]int) [][2]int {
    ans := y
    for i := range x {
        if !isIn(x[i], y) {
            ans = append(ans, x[i]) 
        }
    }
    return ans
}

func isIn(x [2]int, y [][2]int) bool {
    for _, v := range y {
        if v == x {
            return true
        }
    }
    return false
}

func nbrs(node [2]int,inp []string) [][2]int {
    var nbr [][2]int
    if node[0] >= 1 && (string(inp[node[0]-1][node[1]]) == "." || string(inp[node[0]-1][node[1]]) == "S") {
        nbr = append(nbr, [2]int{node[0]-1, node[1]})
    }
    if node[0] + 1 < len(inp) && (string(inp[node[0]+1][node[1]]) == "." || string(inp[node[0]+1][node[1]]) == "S") {
        nbr = append(nbr, [2]int{node[0]+1, node[1]})
    }
    if node[1] >= 1 && (string(inp[node[0]][node[1]-1]) == "." || string(inp[node[0]][node[1]-1]) == "S") {
        nbr = append(nbr, [2]int{node[0], node[1]-1})
    }
    if node[1] + 1 < len(inp[node[0]]) && (string(inp[node[0]][node[1]+1]) == "." || string(inp[node[0]][node[1]+1]) == "S") {
        nbr = append(nbr, [2]int{node[0], node[1]+1})
    }
    //fmt.Println(nbr)
    return nbr
}
