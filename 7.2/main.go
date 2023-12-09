package main

import (
	"os"
	"strings"
    "strconv"
    "fmt"
)

func main() {

    parseData, _ := os.ReadFile("../7.1/input.7.1.txt")
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

    var ans int64 = 0
    for key, val := range cardMaps {
        r := getRank(key, cardMaps)
        //fmt.Println("Rank of", key, " = ", r)
        ans = ans + int64(r) * int64(val)
    }
    fmt.Println(ans)
}

func getMaxVal(m map[string]int) string {
    // make sure not to count J it was headache to debug
    ans := -1
    s := "J" // to handle JJJJJ case
    for key, val := range m {
        if key != "J" && val > ans {
            ans = val
            s = key
        }
    }
    return s
}

func getRank(s string, inp map[string]int) int {
    
    rank := 1
    for key := range inp {
        //fmt.Println("Is ", s, " > ", key, " ",isGreater(s, key))
        if isGreater(s, key) {
            //fmt.Println("I am incrementing rank")
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
        "T": 10,
        "9": 9,
        "8": 8,
        "7": 7,
        "6": 6,
        "5": 5,
        "4": 4,
        "3": 3,
        "2": 2,
        "J": 1,
    } 
    copyFirst := strings.ReplaceAll(first, "J", getMaxVal(getCardMap(first)))
    copySecond := strings.ReplaceAll(second, "J", getMaxVal(getCardMap(second)))
    f, s := getCardMap(copyFirst), getCardMap(copySecond)   

    if len(f) != len(s) {
        return len(f) < len(s)
    }
    // below length of f & s are equal

    if typeClassifier(copyFirst) == typeClassifier(copySecond) {
        for i := 0; i < len(first); i++ {
            if relativeStrength[first[i:i+1]] != relativeStrength[second[i:i+1]] {
                return relativeStrength[first[i:i+1]] > relativeStrength[second[i:i+1]] 
            }
        }
    }
        
    return typeClassifier(copyFirst) > typeClassifier(copySecond)
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
