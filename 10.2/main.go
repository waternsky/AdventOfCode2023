package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
    
    parseData, _ := os.ReadFile("../10.1/input.10.1.txt")
    parseAsStr := string(parseData)
    // make sure to change r & l in cycle
    /*
    parseAsStr = `
FF7FSF7F7F7F7F7F---7
L|LJ||||||||||||F--J
FL-7LJLJ||||||LJL-77
F--JF--7||LJLJ7F7FJ-
L---JF-JLJ.||-FJLJJ7
|F|F-JF---7F7-L7L|7|
|FFJF7L7F-JF7|JL---7
7-L-JL7||F7|L7F-7F7|
L.L7LFJ|||||FJL7||LJ
L7JLJL-JLJLJL--JLJ.L
    `
    */
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

    //fmt.Println(sIndex)

    //fmt.Println(string(splitData[sIndex[0]-1][sIndex[1]])) // -
    //fmt.Println(string(splitData[sIndex[0]][sIndex[1]-1])) // -
    //fmt.Println(string(splitData[sIndex[0]][sIndex[1]+1])) // L
    //fmt.Println(string(splitData[sIndex[0]+1][sIndex[1]])) // |
    // This shows S is 7

    loop := cycle(splitData, sIndex) // note my loop is counter-clockwise helps in counting
    fmt.Println(loop)
    nbrCount := 0
    for i := 0; i < len(splitData); i++ {
        tmp := make([][2]int, 0, len(splitData[i]))
        for j := 0; j < len(loop); j++ {
            if loop[j][0] == i {
                tmp = append(tmp, loop[j])
            }
        }
        // sort tmp
        for p := 0; p < len(tmp); p++ {
            for q := 0; q < len(tmp) - p - 1; q++ {
                if tmp[q][1] > tmp[q+1][1] {
                    temp := tmp[q]
                    tmp[q] = tmp[q+1]
                    tmp[q+1] = temp
                }
            }
        }

        tmp2 := make([][3]int, 0, len(tmp))
        for j := 0; j < len(tmp); j++ {
            for k := 0; k < len(loop); k++ {
                if loop[k] == tmp[j] {
                    next := loop[(k+1)%len(loop)]
                    prev := loop[(len(loop)+k-1)%len(loop)]
                    if next[0] == tmp[j][0] && prev[0] == tmp[j][0] {
                        tmp2 = append(tmp2, [3]int{tmp[j][0], tmp[j][1], 0})
                    } else {
                        if next[0] > tmp[j][0] || tmp[j][0] > prev[0] {
                            tmp2 = append(tmp2, [3]int{tmp[j][0], tmp[j][1], -1})
                        }
                        if next[0] < tmp[j][0] || tmp[j][0] < prev[0] {
                            tmp2 = append(tmp2, [3]int{tmp[j][0], tmp[j][1], 1})
                        }
                    }
                }
            }
        }
        //fmt.Println(tmp2)
        
        for j := 0; j < len(tmp2); j++ {
            if tmp2[0][2] == -1 && tmp2[j][2] == -1 {
                if j < len(tmp2) - 1 && tmp2[j+1][2] == 1 {
                    nbrCount += tmp2[j+1][1] - tmp2[j][1] - 1
                    if tmp2[j+1][1] != tmp2[j][1] + 1 {
                        fmt.Println(tmp2[j], tmp2[j+1])
                    }
                    j=j+1
                    continue
                }
            }
            if tmp2[0][2] == 1 && tmp2[j][2] == 1 {
                if j < len(tmp2) - 1 && tmp2[j+1][2] == -1 {
                    nbrCount += tmp2[j+1][1] - tmp2[j][1] - 1
                    if tmp2[j+1][1] != tmp2[j][1] + 1 {
                        fmt.Println(tmp2[j], tmp2[j+1])
                    }
                    j=j+1
                    continue
                }
            }
        }
    }

    fmt.Println(nbrCount)
}

func cycle(grid []string, start [2]int) [][2]int {
    prevr, prevl := start, start
    r := [2]int{start[0] + 1, start[1]} // main input
    l := [2]int{start[0], start[1] - 1} // main input
    count := 1 // farthest point from start
    var loop [][2]int = [][2]int{l, start, r}
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
        loop = append(loop, r)
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
        loop = append([][2]int{l}, loop...)
        count++
    }
    //fmt.Println(count)

    return loop[1:] // to not count farthest point twice
}
