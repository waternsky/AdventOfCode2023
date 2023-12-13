package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
    
    parseData, _ := os.ReadFile("./input.12.1.txt")
    parseAsStr := string(parseData)
    /*
    parseAsStr = `
???.### 1,1,3
.??..??...?##. 1,1,3
?#?#?#?#?#?#?#? 1,3,1,6
????.#...#... 4,1,1
????.######..#####. 1,6,5
?###???????? 3,2,1
    `
    */
    splitData := strings.Split(strings.TrimSpace(parseAsStr), "\n")

    total := 0
    for i := range splitData {
        line := splitData[i]
        s := strings.TrimSpace(strings.Split(line, " ")[0])
        arr := strings.TrimSpace(strings.Split(line, " ")[1])
        nums := make([]int, 0, len(arr))
        for _, v := range strings.Split(arr, ",") {
            integer, _ := strconv.ParseInt(v, 10, 64)
            nums = append(nums, int(integer))
        }
        fmt.Println(s, nums, springFind(s, nums))
        total += springFind(s, nums)
    }
    fmt.Println(total)
}

func springFind(springs string, nums []int) int {
    re := regexp.MustCompile("#")
    if springs == "" {
        if nums == nil {
            return 1
        }
        return 0
    }
    if nums == nil {
        if re.FindString(springs) == "#" {
            return 0
        }
        return 1
    }
    
    count := 0
    if string(springs[0]) == "." || string(springs[0]) == "?" {
        count += springFind(springs[1:], nums)
    }
    if string(springs[0]) == "?" || string(springs[0]) == "#" {
        if nums[0] <= len(springs) {
            d := true
            for i := range springs[:nums[0]] {
                d = d && string(springs[i]) != "."
                //if string(springs[i]) == "." {
                //    d = false
                //}
            }
            if d && (nums[0] == len(springs) || string(springs[nums[0]]) != "#") {
                if nums[0] == len(springs) {
                    if len(nums) == 1 {
                        count += springFind("", nil)
                    } else {
                        count += springFind("", nums[1:])
                    }
                } else {
                    if len(nums) ==  1 {
                        count += springFind(springs[nums[0]+1:], nil)
                    } else {
                        count += springFind(springs[nums[0]+1:], nums[1:])
                    }
                }
            }
        }
    }

    return count
}
