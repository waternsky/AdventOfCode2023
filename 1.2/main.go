package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 55258
func main() {

    parsedData, _ := os.ReadFile("../1.1/input.1.1.txt")

    parsedDataStr := strings.TrimSpace(string(parsedData))

    splitGames := strings.Split(parsedDataStr, "\n")

    fmt.Println(splitGames)

    replace := map[string]string{
        "one": "1",
        "two": "2",
        "three": "3",
        "four": "4",
        "five": "5",
        "six": "6",
        "seven": "7",
        "eight": "8",
        "nine": "9",
    }
    _, values := keyValues(replace)
    var sum int64 = 0
    for idx := range splitGames {
        splitGames[idx] = convert(splitGames[idx], replace)
        sum += int64(getNum(splitGames[idx], values))
    }

    fmt.Println(splitGames, sum)
    fmt.Println(getNum(convert("twone", replace), values))
    fmt.Println(convert("twone", replace))

    fmt.Println(getNum(convert("7nineight", replace), values))
    fmt.Println(convert("7nineight", replace))

    first, last := 0, 0
    var sumo int64 = 0
    dFound := false
    
    final := strings.Join(splitGames, "\n") + "\n"
    for _, ch := range final {
        chs := string(ch)
        i, err := strconv.ParseInt(chs, 10, 64) 

        if chs == "\n" {
            sumo += int64(last + 10 * first)
            first = 0
            last = 0
            dFound = false
        }

        if err != nil {
            continue
        } 

        if dFound {
            last = int(i)
        } else {
            first = int(i)
            last = int(i)
            dFound = true
        }
    }

    fmt.Println(sumo)

}

func getNum(st string, arr []string) int {
    
    first, last := 0, 0
    out1:
    for i := 0; i < len(st); i++ {
        for j := 0; j < len(arr); j++ {
            if string(st[i]) == arr[j] {
                num1, _ := strconv.ParseInt(arr[j], 10, 64)
                first = int(num1)
                break out1
            }
        }
    }
    out2:
    for i := 0; i < len(st); i++ {
        for j := 0; j < len(arr); j++ {
            if string(st[len(st) - i - 1]) == arr[j] {
                num2, _ := strconv.ParseInt(arr[j], 10, 64)
                last = int(num2)
                break out2
            }
        }
    }
    return last + 10 * first
}

func keyValues(m map[string]string) ([]string, []string) {
    keys := make([]string, 0, len(m))
    values := make([] string, 0, len(m))

    for key, value := range m {
        keys = append(keys, key)
        values = append(values, value)
    }

    return keys, values
}

func sortByLenAsc(arr []string, asc bool) {
    if asc {
        for i := 0; i < len(arr); i++ {
            for j := 0; j < len(arr) - i - 1; j++ {
                if len(arr[j]) > len(arr[j+1]) {
                    tmp := arr[j]
                    arr[j] = arr[j+1]
                    arr[j+1] = tmp
                }         
            }
        }
    } else {
        for i := 0; i < len(arr); i++ {
            for j := 0; j < len(arr) - i - 1; j++ {
                if len(arr[j]) < len(arr[j+1]) {
                    tmp := arr[j]
                    arr[j] = arr[j+1]
                    arr[j+1] = tmp
                }         
            }
        }
    }
}

func convert(str string, replace map[string]string) string {
    keys, _ := keyValues(replace)
    sortByLenAsc(keys, true)    

    out1:
    for i := 0; i < len(str) - len(keys[0]); i++ {
        if string(str[i]) == "1" || string(str[i]) == "2" || string(str[i]) == "3" || string(str[i]) == "4" || string(str[i]) == "5" || string(str[i]) == "6" || string(str[i]) == "7" || string(str[i]) == "8" || string(str[i]) == "9" {
            break out1
        }
        for j := 0; j < len(keys); j++ {
            if i < len(str) - len(keys[j]) && str[i:i+len(keys[j])] == keys[j] {
                str = str[0:i]+replace[keys[j]]+str[i+len(keys[j]):]
                break out1
            }
        } 
    }
    sortByLenAsc(keys, false)
    out2:
    for i := len(str); i > len(keys[len(keys)-1]); i-- {
        if string(str[i-1]) == "1" || string(str[i-1]) == "2" || string(str[i-1]) == "3" || string(str[i-1]) == "4" || string(str[i-1]) == "5" || string(str[i-1]) == "6" || string(str[i-1]) == "7" || string(str[i-1]) == "8" || string(str[i-1]) == "9" {
            break out2
        }
        for j := 0; j < len(keys); j++ {
            if i >= len(keys[j]) && str[i-len(keys[j]):i] == keys[j] {
                str = str[0:i-len(keys[j])]+replace[keys[j]]+str[i:]
                break out2
            }
        } 
    }

    return str
} 
