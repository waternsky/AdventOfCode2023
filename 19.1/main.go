package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
    parseData, _ := os.ReadFile("./input.19.1.txt")
    parseAsStr := string(parseData)

    parseAsStr = strings.TrimSpace(parseAsStr)
    splitData := strings.Split(parseAsStr, "\n\n")

    workflows := splitData[0]
    parts := splitData[1]

    ans := 0
    for _ , v := range strings.Split(parts, "\n") {
        part := ratings(v)
        x := partThroughWorkflows(part, workflows)
        if x == "A" {
            for _, num := range part {
                ans += num
            }
        }
    }

    fmt.Println(ans)
}

func checkRule(part map[string]int, ruleArr []string) string {
    re := regexp.MustCompile("[<=>:]")
    ans := ruleArr[len(ruleArr)-1]
    for i := 0; i < len(ruleArr) - 1; i++ {
        rule := ruleArr[i]
        tmp := re.Split(rule, -1)
        integer, _ := strconv.ParseInt(tmp[1], 10, 64)
        if (re.FindString(rule) == ">" && part[tmp[0]] > int(integer)) ||
        (re.FindString(rule) == "<" && part[tmp[0]] < int(integer)) ||
        (re.FindString(rule) == "=" && part[tmp[0]] == int(integer)) {
            ans = tmp[2]
            break
        }
    }
    return ans
}

func partThroughWorkflows(part map[string]int, workflows string) string {
    m := make(map[string][]string)
    for _, w := range strings.Split(workflows, "\n") {
        name, ruleArr := parseWorkflow(w)
        m[name] = ruleArr
    }
    x := "in"
    for x != "A" && x != "R" {
        x = checkRule(part, m[x])
    } 
    return x
}

func ratings(inp string) map[string]int {
    rating := inp[1:len(inp)-1]
    arr := strings.Split(rating, ",")
    m := make(map[string]int)
    for i := range arr {
        tmp := strings.Split(arr[i], "=")
        integer, _ := strconv.ParseInt(tmp[1], 10, 64)
        m[tmp[0]] = int(integer)
    }
    return m
}

func parseWorkflow(inp string) (string, []string) {
    tmp := strings.Split(inp, "{")
    name := tmp[0]
    rules := tmp[1][:len(tmp[1])-1]
    ruleArr := strings.Split(rules, ",")
    return name, ruleArr
}
