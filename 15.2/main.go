package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
    parseData, _ := os.ReadFile("../15.1/input.15.1.txt")
    parseAsStr := string(parseData)

    splitData := strings.Split(strings.TrimSpace(parseAsStr), ",")

    m := make(map[int][]lens)
    for i := range splitData {
        st := splitData[i]
        lbf := strings.Split(st, "=")
        if len(lbf) == 2 {
            integer, _ := strconv.ParseInt(lbf[1], 10, 64)
            l := lens {
                label: lbf[0],
                focal_length: int(integer),
            }
            boxNum := algo(lbf[0])
            if isIn(l, m[boxNum]) {
                for j := range m[boxNum] {
                    if m[boxNum][j].label == l.label {
                        m[boxNum][j].focal_length = l.focal_length
                        break
                    }
                }
            } else {
                m[boxNum] = append(m[boxNum], l)
            }
        } else {
            label := strings.Split(st, "-")[0]
            boxNum := algo(label)
            tmp := lens {
                label: label,
                focal_length: -1,
            }
            if isIn(tmp, m[boxNum]) {
                idx := 0
                in:
                for j := range m[boxNum] {
                    if m[boxNum][j].label == tmp.label {
                        idx = j
                        break in
                    }
                }
                m[boxNum] = append(m[boxNum][:idx], m[boxNum][idx+1:]...)
            }
        }
    }

    ans := 0
    for boxNum, lenses := range m {
        for i := range lenses {
            ans += (boxNum+1) * (i+1) * lenses[i].focal_length
        } 
    }
    fmt.Println(ans)
}

func isIn(l lens, arr []lens) bool {
    for i := range arr {
        if l.label == arr[i].label {
            return true
        }
    }
    return false
}

type lens struct {
    label string
    focal_length int
}

func algo(s string) int {
    val := 0
    for i := range s {
        ascii := int(s[i]) // rune 
        val += ascii
        val *= 17
        val = val % 256
    }
    return val
}
