package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {

    parseData, _ := os.ReadFile("../3.1/input.3.1.txt")

    parseAsStr := strings.TrimSpace(string(parseData))

    splitData := strings.Split(parseAsStr, "\n")

    fmt.Println(gearRatio(splitData))
}

func checkStarIdx(inp string) [][]int {
    re := regexp.MustCompile("\\*")
    return re.FindAllStringIndex(inp, -1)
} 

func getNumIdx(inp string) [][]int {
    re := regexp.MustCompile("\\d+")
    return re.FindAllStringIndex(inp, -1)
}

func gearRatio(inp []string) int64 {

    var sum int64 = 0
    for i := range inp {
        stars := checkStarIdx(inp[i])
        nums := getNumIdx(inp[i])
        for j := range stars {
            star := stars[j]
            m := make(map[int][][]int)
            // inline checks
            for k := range nums {
                if star[0] == nums[k][1] {
                    m[i] = append(m[i], nums[k]) 
                }
                if star[1] == nums[k][0] {
                    m[i] = append(m[i], nums[k])
                }
            }
            // aboveline checks
            if i > 0 {
                numsAbove := getNumIdx(inp[i-1])
                for k := range numsAbove {
                    if numsAbove[k][1] == star[0] {
                        m[i-1] = append(m[i-1], numsAbove[k])
                    }
                    if numsAbove[k][0] == star[1] {
                        m[i-1] = append(m[i-1], numsAbove[k])
                    }
                    if numsAbove[k][0] <= star[0] && numsAbove[k][1] >= star[1] {
                        m[i-1] = append(m[i-1], numsAbove[k])
                    } 
                }
            }
            // belowline checks
            if i < len(inp) - 1 {
                numsBelow := getNumIdx(inp[i+1])
                for k := range numsBelow {
                    if numsBelow[k][1] == star[0] {
                        m[i+1] = append(m[i+1], numsBelow[k])
                    }
                    if numsBelow[k][0] == star[1] {
                        m[i+1] = append(m[i+1], numsBelow[k])
                    }
                    if numsBelow[k][0] <= star[0] && numsBelow[k][1] >= star[1] {
                        m[i+1] = append(m[i+1], numsBelow[k])
                    } 
                }
            }
            
            tC := 0
            for _, val := range m {
                tC += len(val)
            }
            if tC == 2 {
                var gear int64 = 1
                for keyI, valI := range m {
                    for numIdx := range valI {
                        integer, _ := strconv.ParseInt(inp[keyI][valI[numIdx][0]:valI[numIdx][1]] , 10, 64)
                        gear *= integer
                    }
                }
                sum += gear
            }
        }
    }

    return sum
}
