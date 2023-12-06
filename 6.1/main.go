package main

import (
    "fmt"
)

func main() {
    /*
    inp := `
    Time:        40     81     77     72
    Distance:   219   1012   1365   1089
    `
    */
    fmt.Println(getTimeRange(40, 219) * getTimeRange(81, 1012) * getTimeRange(77, 1365) * getTimeRange(72, 1089))
    fmt.Println(getTimeRange(40817772, 219101213651089))
}

func getTimeRange(time int, distance int) int {
    
    distances := make([]int, time+1, time+1)
    for i := 0; i <= time; i++ {
        distances[i] = (time - i) * i
    }
    winTime := make([]int, 0, time+1)
    for i := 0; i < len(distances); i++ {
        if (distances[i] > distance) {
            winTime = append(winTime, i)
        }
    }
    
    return winTime[len(winTime)-1] - winTime[0] + 1
}
