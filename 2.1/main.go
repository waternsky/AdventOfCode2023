package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
    
    parsedData, _ := os.ReadFile("./input.2.1.txt")

    parsedDataStr := strings.TrimSpace(string(parsedData))

    splitGames := strings.Split(parsedDataStr, "\n")

    sum := 0
    for idx := 0; idx < len(splitGames); idx++ {
        games := splitGames[idx]
        id, validGame := isValidGame(games, 12, 14, 13)
        if validGame {
            sum += id
        }
    }
    fmt.Println(sum)
}

func isValidGame(t string, rC int, bC int, gC int) (int, bool) {
    
    parse := strings.Split(t, " ")

    gameId, _ := strconv.ParseInt(parse[1][0:(len(parse[1])-1)], 10, 64)
    id := int(gameId)

    splG := parse[2:]
    games := strings.Join(splG, " ")
    game := strings.Split(games, ";")

    vG := true
    for i := 0; i < len(game); i++ {
        gm := game[i]
        _, _, _, cG := rbg(gm, rC, bC, gC)
        vG = vG && cG
    }

    return id, vG
}

func rbg(game string, rC int, bC int, gC int) (int, int, int, bool) {
    
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


    return r, b, g, r <= rC && b <= bC && g <= gC
}
