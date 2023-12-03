package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {

    parseData, _ := os.ReadFile("./input.3.1.txt")

    parseAsStr := string(parseData)
    //parseAsStr = "467..114..\n...*......\n..35..633.\n......#...\n617*......\n.....+.58.\n..592.....\n......755.\n...$.*....\n.664.598.."

    splitData := strings.Split(strings.TrimSpace(parseAsStr), "\n")

    //fmt.Println(splitData[0])

    //p := splitData[0]

    //fmt.Println(strings.Split(p, ""))

    fmt.Println(sumAll(splitData)) // incorrect sum here 
    fmt.Println(sumAll2(splitData))
    fmt.Println(sumByDots(splitData))
    //fmt.Println(numInstances("..&%467..114.."))
    //fmt.Println(getNthNumIdx("..&%467..114..44", 4))

    //fmt.Println(getNumIdx("abcd6"))

    //fmt.Println(p[len(p):])
    //fmt.Println(checkNbrDots(splitData, 0, 5, 7))
}

func contains(elm string, arr []string) bool {
    ans := false
    for idx := range arr {
        if arr[idx] == elm {
            ans = true
            break
        }
    }
    return ans
}

func numInstances(line string) int {

    /*
    re := regexp.MustCompile("[^0-9]+")


    split := re.Split(line, -1)

    js := strings.Join(split, " ")
    return len(strings.Split(strings.TrimSpace(js), " "))
    */

    re := regexp.MustCompile("[0-9]+")
    return len(re.FindAllString(line, -1))
}

func sumAll2(inp []string) int64 {
    var sum int64 = 0
    for idx := range inp {
        line := inp[idx]
        re := regexp.MustCompile("[0-9]+")
        arr := re.FindAllString(line, -1)
        for i := range arr {
            integer, _ := strconv.ParseInt(arr[i], 10, 64)
            sum += integer
        }
    }
    return sum
}

func getNthNumIdx(line string, nth int) (int, int) {

    if nth <= 0 {
        return -1, -1
    }
    copyLine := line
    nb := numInstances(line)
    idx1, idx2 := getFirstNumIdx(copyLine)
    tmp := 0
    for i := 1; i < nth; i++ {
        if nth > nb {
            idx1 = -1
            idx2 = -1
            break
        }
        tmp = len(line[:idx2+1])
        copyLine = line[idx2+1:]
        idx1, idx2 = getFirstNumIdx(copyLine)
        idx1 += tmp
        idx2 += tmp
    }
    return idx1, idx2
}

func getFirstNumIdx(line string) (int, int) {
    nums := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
    idx1, idx2 := -1, -1
    dprev := false
    for idx := range line {
        ch := string(line[idx])
        if contains(ch, nums) {
            if dprev {
                idx2 = idx
                continue
            }
            dprev = true
            idx1 = idx
            idx2 = idx
        } else {
            if dprev {
                break
            }
            dprev = false
        }
            
    }
    return idx1, idx2
}

func checkNbrDots(inp []string, lineIdx int, idx1 int, idx2 int) bool {
    line := inp[lineIdx]

    if idx1 > idx2 {
        return false
    }

    if lineIdx == 0 {
        ans := true
        for i := idx1; i <= idx2; i++ {
            //fmt.Println(inp[lineIdx+1][i])
            ans = ans && string(inp[lineIdx+1][i]) == "."
        }
        if idx1 == 0 {
            if idx2 == len(line) - 1 {
                return ans
            }
            ans = ans && string(inp[lineIdx+1][idx2+1]) == "." && string(line[idx2+1]) == "."
            return ans
        }
        if idx2 == len(line) - 1 {
            if idx1 == 0 {
                return ans
            }
            ans = ans && string(inp[lineIdx+1][idx1-1]) == "." && string(line[idx1-1]) == "."
            return ans
        }
        //fmt.Println(string(line[idx1-1]))
        //fmt.Println(string(line[idx2+1]))
        //fmt.Println(string(inp[lineIdx+1][idx1-1]))
        //fmt.Println(string(inp[lineIdx+1][idx2+1]))
        ans = ans && string(line[idx1-1]) == "." && string(line[idx2+1]) == "." && string(inp[lineIdx+1][idx1-1]) == "." && string(inp[lineIdx+1][idx2+1]) == "."
        return ans
    }

    if lineIdx == len(inp) - 1 {
        ans := true
        for i := idx1; i <= idx2; i++ {
            ans = ans && string(inp[lineIdx-1][i]) == "."
        }
        if idx1 == 0 {
            if idx2 == len(line) - 1 {
                return ans
            }
            ans = ans && string(inp[lineIdx-1][idx2+1]) == "." && string(line[idx2+1]) == "."
            return ans
        }
        if idx2 == len(line) - 1 {
            if idx1 == 0 {
                return ans
            }
            ans = ans && string(inp[lineIdx-1][idx1-1]) == "." && string(line[idx1-1]) == "."
            return ans
        }
        ans = ans && string(line[idx1-1]) == "." && string(line[idx2+1]) == "." && string(inp[lineIdx-1][idx1-1]) == "." && string(inp[lineIdx-1][idx2+1]) == "."
        return ans
    }

    ans := true
    for i := idx1; i <= idx2; i++ {
        ans = ans && string(inp[lineIdx-1][i]) == "."
        ans = ans && string(inp[lineIdx+1][i]) == "."
    }
    if idx1 == 0 {
        if idx2 == len(line) - 1 {
            return ans
        }
        ans = ans && string(inp[lineIdx-1][idx2+1]) == "." && string(line[idx2+1]) == "." && string(inp[lineIdx+1][idx2+1]) == "."
        return ans
    }
    if idx2 == len(line) - 1 {
        if idx1 == 0 {
            return ans
        }
        ans = ans && string(inp[lineIdx-1][idx1-1]) == "." && string(line[idx1-1]) == "." && string(inp[lineIdx+1][idx1-1]) == "."
        return ans
    }
    ans = ans && string(line[idx1-1]) == "." && string(line[idx2+1]) == "." && string(inp[lineIdx-1][idx1-1]) == "." && string(inp[lineIdx-1][idx2+1]) == "." && string(inp[lineIdx+1][idx1-1]) == "." && string(inp[lineIdx+1][idx2+1]) == "."
    return ans
}

func sumByDots(inp []string) int64 {
    var sum int64 = 0
    for idx := range inp {
        //fmt.Println("I start")
        line := inp[idx]
        i := 1
        idx1, idx2 := getNthNumIdx(line, i)
        //fmt.Println("first: ", idx1, idx2)
        for idx1 > -1 && idx2 > -1 {
            //fmt.Println(checkNbrDots(inp, idx, idx1, idx2))
            if (checkNbrDots(inp, idx, idx1, idx2)) {
                //fmt.Println("iaminevitable")
                integer, _ := strconv.ParseInt(line[idx1:idx2+1], 10, 64)
                sum += integer
            }
            i++
            idx1, idx2 = getNthNumIdx(line, i)
            //fmt.Println(idx1, idx2)
        }

    }
    return sum
}

func sumAll(inp []string) int64 {
    var sum int64 = 0
    for i := 0; i < len(inp); i++ {
        line := inp[i]
        dprev := false
        var num int64 = 0
        in1:
        for j := 0; j < len(line); j++ {
            in, err := strconv.ParseInt(string(line[j]), 10, 64)
            if err == nil {
                if dprev {
                    num = in + 10 * num
                    continue in1
                }
                num = in
                dprev = true
            } else {
                sum += num
                dprev = false
                num = 0
            }
        }
    }
    return sum
}
