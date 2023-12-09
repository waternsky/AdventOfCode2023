package main

import (
	"os"
	"strings"
    "fmt"
)

func gcd(a int, b int) int {
    if b == 0 {
        return a
    }
    return gcd(b, a % b)
}

func lcm(arr []int) int64 {
    lcm := int64(arr[0])
    for _, num := range arr[1:] {
        lcm = lcm * int64(num) / int64(gcd(int(lcm), num))
    }
    return lcm
}

func main() {
    parseData, _ := os.ReadFile("../8.1/input.8.1.txt")
    parseAsStr := string(parseData)
    /*
    parseAsStr = `
    LR

11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)
    `
    */
    splitData := strings.Split(strings.TrimSpace(parseAsStr), "\n\n")
    instructions := strings.TrimSpace(splitData[0])
    //fmt.Println(instructions)

    node := strings.Split(strings.TrimSpace(splitData[1]), "\n")
    m := make(map[string][]string)
    for i := range node {
        key := strings.TrimSpace(strings.Split(node[i], "=")[0]) 
        val := strings.TrimSpace(strings.Split(node[i], "=")[1]) 
        li := make([]string, 0, 2)
        li = append(li, strings.Split(val, ",")[0][1:])
        li = append(li, strings.TrimSpace(strings.Split(val, ",")[1])[:3])
        m[key] = li
    }
    //fmt.Println(m)
    x := make([]string, 0, 6)
    for key := range m {
        if string(key[2]) == "A" {
            x = append(x, key)
        }
    }
    fmt.Println(x)
   
    for i := 0; i < len(x); i++ {
        s := x[i]
        count := 0
        zcount := 0
        zpositions := make(map[string][]int)
        for true {
            if zcount >= 1 {
                break
            }
            if string(s[2]) == "Z" {
                zpositions[s] = append(zpositions[s], count)
                zcount++
            }
            if string(instructions[count % len(instructions)]) == "L" {
                s = m[s][0]
            }
            if string(instructions[count % len(instructions)]) == "R" {
                s = m[s][1]
            }
            count++
        }
        fmt.Println(zpositions)
    }
    fmt.Println(lcm([]int{19667, 21883, 11911, 13019, 14681, 16897}))

    /*
    This solution takes too long to run (i let it ran for over 90 mins with no luck) as answer is of order ~ 10^13
    var count int64 = 0
    for !endWithZ(x) {
        ins := string(instructions[int(count % int64(len(instructions)))])
        if ins == "L" {
            for i := range x {
                x[i] = m[x[i]][0]
            }
        }
        if ins == "R" {
            for i := range x {
                x[i] = m[x[i]][1]
            }
        }
        count++
    }
    fmt.Println(count)
    */
}

func endWithZ(arr []string) bool {
    ans := true
    for i := range arr {
        if string(arr[i][2]) != "Z" {
            ans = false
            break
        }
    } 
    return ans
}
