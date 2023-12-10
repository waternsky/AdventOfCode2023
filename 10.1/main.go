package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
    
    parseData, _ := os.ReadFile("./input.10.1.txt")
    parseAsStr := string(parseData)

    splitData := strings.Split(strings.TrimSpace(parseAsStr), "\n")

    sIndex := [2]int{-1, -1}
    out:
    for i := 0; i < len(splitData); i++ {
        for j := 0; j < len(splitData[i]); j++ {
            if string(splitData[i][j]) == "S" {
                sIndex = [2]int{i, j}
                break out
            }
        }
    }

    //fmt.Println(sString, sIndex)

    //fmt.Println(string(splitData[sIndex[0]-1][sIndex[1]])) // -
    //fmt.Println(string(splitData[sIndex[0]][sIndex[1]-1])) // -
    //fmt.Println(string(splitData[sIndex[0]][sIndex[1]+1])) // L
    //fmt.Println(string(splitData[sIndex[0]+1][sIndex[1]])) // |
    // This shows S is 7

    walk(splitData, sIndex)
}

func walk(grid []string, start [2]int) {
    prevr, prevl := start, start
    // In current case S is 7 hence l & r defined as below
    r := [2]int{start[0] + 1, start[1]}
    l := [2]int{start[0], start[1] - 1}
    count := 1
    for r != l {
        switch piper := string(grid[r[0]][r[1]]); piper {
            case "|":
                if prevr == [2]int{r[0]-1, r[1]} {
                    prevr = r
                    r = [2]int{r[0]+1, r[1]} 
                } else {
                    prevr = r
                    r = [2]int{r[0]-1, r[1]}
                }
            case "-":
                if prevr == [2]int{r[0], r[1]-1} {
                    prevr = r
                    r = [2]int{r[0], r[1]+1} 
                } else {
                    prevr = r
                    r = [2]int{r[0], r[1]-1}
                }
            case "L":
                if prevr == [2]int{r[0]-1, r[1]} {
                    prevr = r
                    r = [2]int{r[0], r[1]+1} 
                } else {
                    prevr = r
                    r = [2]int{r[0]-1, r[1]}
                }
            case "J":
                if prevr == [2]int{r[0]-1, r[1]} {
                    prevr = r
                    r = [2]int{r[0], r[1]-1} 
                } else {
                    prevr = r
                    r = [2]int{r[0]-1, r[1]}
                }
            case "7":
                if prevr == [2]int{r[0]+1, r[1]} {
                    prevr = r
                    r = [2]int{r[0], r[1]-1} 
                } else {
                    prevr = r
                    r = [2]int{r[0]+1, r[1]}
                }
            case "F":
                if prevr == [2]int{r[0]+1, r[1]} {
                    prevr = r
                    r = [2]int{r[0], r[1]+1} 
                } else {
                    prevr = r
                    r = [2]int{r[0]+1, r[1]}
                }
        }
        switch pipel := string(grid[l[0]][l[1]]); pipel {
            case "|":
                if prevl == [2]int{l[0]-1, l[1]} {
                    prevl = l
                    l = [2]int{l[0]+1, l[1]} 
                } else {
                    prevl = l
                    l = [2]int{l[0]-1, l[1]}
                }
            case "-":
                if prevl == [2]int{l[0], l[1]-1} {
                    prevl = l
                    l = [2]int{l[0], l[1]+1} 
                } else {
                    prevl = l
                    l = [2]int{l[0], l[1]-1}
                }
            case "L":
                if prevl == [2]int{l[0]-1, l[1]} {
                    prevl = l
                    l = [2]int{l[0], l[1]+1} 
                } else {
                    prevl = l
                    l = [2]int{l[0]-1, l[1]}
                }
            case "J":
                if prevl == [2]int{l[0]-1, l[1]} {
                    prevl = l
                    l = [2]int{l[0], l[1]-1} 
                } else {
                    prevl = l
                    l = [2]int{l[0]-1, l[1]}
                }
            case "7":
                if prevl == [2]int{l[0]+1, l[1]} {
                    prevl = l
                    l = [2]int{l[0], l[1]-1} 
                } else {
                    prevl = l
                    l = [2]int{l[0]+1, l[1]}
                }
            case "F":
                if prevl == [2]int{l[0]+1, l[1]} {
                    prevl = l
                    l = [2]int{l[0], l[1]+1} 
                } else {
                    prevl = l
                    l = [2]int{l[0]+1, l[1]}
                }
        }
        count++
    }
    fmt.Println(count)
}
