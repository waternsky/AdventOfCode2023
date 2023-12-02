package main

import (
    "fmt"
    "os"
    "strings"
    "strconv"
)

func main() {
    
    parsedData, _ := os.ReadFile("../2.1/input.2.1.txt")

    parsedDataStr := strings.TrimSpace(string(parsedData))

    splitGames := strings.Split(parsedDataStr, "\n")
    
    var sum int64 = 0

    for idx := range splitGames {
        sum += power(splitGames[idx])
    }

    fmt.Println(sum)
}

func power(t string) int64 {
    
    parse := strings.Split(t, " ")

    splG := parse[2:]
    games := strings.Join(splG, " ")
    game := strings.Split(games, ";")

    rMax, bMax, gMax := 0, 0, 0
    for i := 0; i < len(game); i++ {
        gm := game[i]
        r, b, g := rbg(gm)
        rMax = max(rMax, r)
        bMax = max(bMax, b)
        gMax = max(gMax, g)
    }

    return int64(rMax) * int64(bMax) * int64(gMax)
}

func rbg(game string) (int, int, int) {
    
    r, b, g := 0, 0, 0
    
    balls := strings.Split(game, ",")

    for i := 0; i < len(balls); i++ {
        ball := strings.TrimSpace(balls[i])
        numColor := strings.Split(ball, " ")
        if numColor[1] == "green" {
            num,_ := strconv.ParseInt(numColor[0], 10, 64) 
            g += int(num)
        }
        if numColor[1] == "blue" {
            num,_ := strconv.ParseInt(numColor[0], 10, 64) 
            b += int(num)
        }
        if numColor[1] == "red" {
            num,_ := strconv.ParseInt(numColor[0], 10, 64) 
            r += int(num)
        }
    }


    return r, b, g
}
