package main

import (
	"os"
	"strings"
    "fmt"
)

func main() {
    parseData, _ := os.ReadFile("./input.8.1.txt")
    parseAsStr := string(parseData)

    splitData := strings.Split(strings.TrimSpace(parseAsStr), "\n\n")
    instructions := strings.TrimSpace(splitData[0])

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

    x := "AAA"
    count := 0
    for x != "ZZZ" {
        if string(instructions[count % len(instructions)]) == "L" {
            x = m[x][0]
        } 
        if string(instructions[count % len(instructions)]) == "R" {
            x = m[x][1]
        } 
        count++
    }

    fmt.Println(count)
}
