package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Point struct {
    x float64
    y float64
    z float64
}

type Velocity struct {
    vx float64
    vy float64
    vz float64
}

type hailstone struct {
    p Point
    v Velocity
}

func intersection(a hailstone, b hailstone) [2]float64 {
    det := a.v.vx * b.v.vy - a.v.vy * b.v.vx
    if det == 0 {
        return [2]float64{}
    }
    x := b.p.x + (a.v.vy * b.v.vx * (b.p.x - a.p.x) + a.v.vx * b.v.vx * (a.p.y - b.p.y)) / det
    y := a.p.y + (a.v.vy * b.v.vx * (a.p.y - b.p.y) + a.v.vy * b.v.vy * (b.p.x - a.p.x)) / det

    t1 := ((b.p.x - a.p.x) * b.v.vy - (b.p.y - a.p.y) * b.v.vx) / det
    t2 := ((b.p.x - a.p.x) * a.v.vy - (b.p.y - a.p.y) * a.v.vx) / det
    
    if t1 > 0 && t2 > 0 {
        return [2]float64{x, y}
    }
    return [2]float64{} 
}

func check(x [2]float64, param1 float64, param2 float64) bool {
    if len(x) == 0 {
        return false
    }
    if x[0] >= param1 && x[0] <= param2 && x[1] >= param1 && x[1] <= param2 {
        return true
    }
    return false
}

func main() {
    parseData, _ := os.ReadFile("./input.24.1.txt")
    parseAsStr := string(parseData)
    /*
    parseAsStr = `
19, 13, 30 @ -2,  1, -2
18, 19, 22 @ -1, -1, -2
20, 25, 34 @ -2, -2, -4
12, 31, 28 @ -1, -2, -1
20, 19, 15 @  1, -5, -3
    `
    */
    parseAsStr = strings.TrimSpace(parseAsStr)
    splitData := strings.Split(parseAsStr, "\n")

    hailarr := make([]hailstone, 0, len(splitData))
    for i := range splitData {
        pos := strings.Split(splitData[i], "@")[0]
        vel := strings.Split(splitData[i], "@")[1]
        x, _ := strconv.ParseFloat(strings.TrimSpace(strings.Split(pos, ",")[0]), 64)
        y, _ := strconv.ParseFloat(strings.TrimSpace(strings.Split(pos, ",")[1]), 64)
        z, _ := strconv.ParseFloat(strings.TrimSpace(strings.Split(pos, ",")[2]), 64)
        vx, _ := strconv.ParseFloat(strings.TrimSpace(strings.Split(vel, ",")[0]), 64)
        vy, _ := strconv.ParseFloat(strings.TrimSpace(strings.Split(vel, ",")[1]), 64)
        vz, _ := strconv.ParseFloat(strings.TrimSpace(strings.Split(vel, ",")[2]), 64)
        hail := hailstone {
            p: Point {
                x: x,
                y: y,
                z: z,
            },
            v: Velocity {
                vx: vx,
                vy: vy,
                vz: vz,
            },
        }
        hailarr = append(hailarr, hail)
    }

    ans := 0
    for i := 0; i < len(hailarr) - 1; i++ {
        for j := i+1; j < len(hailarr); j++ {
            if check(intersection(hailarr[i], hailarr[j]), 200000000000000, 400000000000000) {
                ans++
            }
        }
    }
    fmt.Println(ans)
}
