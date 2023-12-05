package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
    
    parseData, _ := os.ReadFile("./input.5.1.txt")
    parseAsStr := strings.TrimSpace(strings.TrimSpace(string(parseData)))
    /*
    parseAsStr = strings.TrimSpace(`
seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4
    `)
    */
    splitData := strings.Split(parseAsStr, "\n\n")
    
    re := regexp.MustCompile("\\d+")
    seedStr := re.FindAllString(splitData[0], -1)
    seeds := make([]int64, 0, len(seedStr))
    for _, val := range seedStr {
        integer, _ := strconv.ParseInt(val, 10, 64)
        seeds = append(seeds, integer)
    }

    //fmt.Println("Initial seeds: ", seeds)
    //fmt.Println(convert2(splitData[1], seeds))
    
    ans := seeds
    for i := 1; i < len(splitData); i++ {
        ans = convert2(splitData[i], ans) 
    }
    fmt.Println(minS(ans))
}

func minS(m []int64) int64 {
    ans := m[0]
    for i := range m {
        if m[i] < ans {
            ans = m[i]
        }
    }
    return ans
}

func convert2(inp string, initial []int64) []int64 {
    split := strings.Split(inp, "\n")
    mapFromTo := make([][3]int64, 0, len(split) - 1)
    ans := make([]int64, len(initial) , len(initial))
    re := regexp.MustCompile("\\d+")
    for i := 1; i < len(split); i++ {
        var m [3]int64
        for idx, val := range re.FindAllString(split[i], -1) {
            integer, _ := strconv.ParseInt(val, 10, 64)
            m[idx] = integer
        }
        mapFromTo = append(mapFromTo, m)
    }
    for idx := range initial {
        in:
        for j := range mapFromTo {
            li := mapFromTo[j]
            if li[1] <= initial[idx] && li[1] + li[2] > initial[idx] {
                ans[idx] = li[0] + initial[idx] - li[1]
                break in
            } else {
                ans[idx] = initial[idx]
            } 
        }
    }

    return ans
}

// very inefficient way
func convert(inp string) map[int64]int64 {
    split := strings.Split(inp, "\n")
    re := regexp.MustCompile("\\d+")
    m := make(map[int64]int64)
    for i := 1; i < len(split); i++ {
        inx := make([]int64, 0, 3)
        for _, val := range re.FindAllString(split[i], -1) {
            integer, _ := strconv.ParseInt(val, 10, 64)
            inx = append(inx, integer)
        }
        fmt.Println(inx)
        var tmp int64 = 0
        for tmp < inx[2] {
            m[inx[1] + tmp] = inx[0] + tmp
            tmp++
        }
    }
    return m
}
