package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

    parseData, _ := os.ReadFile("./input.7.1.txt")
    parseAsStr := string(parseData)
    /*
    parseAsStr = `
32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483
    `
    */
    splitData := strings.Split(strings.TrimSpace(parseAsStr), "\n")

    cardMaps := make(map[string]int)

    for i := range splitData {
        in, _ := strconv.ParseInt(strings.Split(splitData[i], " ")[1], 10, 64)
        cardMaps[strings.Split(splitData[i], " ")[0]] = int(in)
    }
    ans := 0
    for key, val := range cardMaps {
        ans += getRank(key, cardMaps) * val
    }
    fmt.Println(ans)
}

func getRank(s string, inp map[string]int) int {
    
    rank := 1
    for key := range inp {
        if s != key && isGreater(s, key) {
            rank++
        }
    }
    return rank 
}

func isGreater(first string, second string) bool {
    relativeStrength := map[string]int{
        "A": 14,
        "K": 13,
        "Q": 12,
        "J": 11,
        "T": 10,
        "9": 9,
        "8": 8,
        "7": 7,
        "6": 6,
        "5": 5,
        "4": 4,
        "3": 3,
        "2": 2,
    } 
    f, s := getCardMap(first), getCardMap(second)    
    if len(f) != len(s) {
        return len(f) < len(s)
    }
    // below length of f & s are equal

    if typeClassifier(first) == typeClassifier(second) {
        for i := 0; i < len(first); i++ {
            if relativeStrength[first[i:i+1]] != relativeStrength[second[i:i+1]] {
                return relativeStrength[first[i:i+1]] > relativeStrength[second[i:i+1]] 
            }
        }
    }
        
    return typeClassifier(first) > typeClassifier(second)
}

func typeClassifier(cards string) int {
    cardMap := getCardMap(cards)
    if len(cards) > 5 || len(cards) < 5 {
        return -1
    }
    if len(cardMap) == 1 {
        return 7 //"Five of a kind"
    }
    if len(cardMap) == 2 {
        for _, val := range cardMap {
            if val == 4 {
                return 6 //"Four of a kind"
            }
            if val == 3 {
                return 5 //"Full house"
            }
        }
    }
    if len(cardMap) == 3 {
        for _, val := range cardMap {
            if val == 3 {
                return 4 //"Three of a kind"
            }
            if val == 2 {
                return 3 //"Two pair"
            }
        }
    }
    if len(cardMap) == 4 {
        return 2 //"One pair"
    }

    return 1 //"High card"
}

func getCardMap(s string) map[string]int {
    m := make(map[string]int)
    for i := range s {
        m[s[i:i+1]]++   
    }
    return m
}
