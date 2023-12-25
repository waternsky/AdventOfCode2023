package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
    
    "gonum.org/v1/gonum/mat"
)

type Point struct {
    x int64
    y int64
    z int64
}

type Velocity struct {
    vx int64
    vy int64
    vz int64
}

type hailstone struct {
    p Point
    v Velocity
}

func coeff(a hailstone, b hailstone) ([2][6]int64, [2]int64) {
    return [2][6]int64{
        {b.v.vy - a.v.vy, -(b.v.vx - a.v.vx), 0, -(b.p.y - a.p.y), b.p.x - a.p.x, 0},
        {b.v.vz - a.v.vz, 0, -(b.v.vx - a.v.vx), -(b.p.z - a.p.z), 0, b.p.x - a.p.x},
    }, [2]int64{a.p.x * a.v.vy - b.p.x * b.v.vy - a.p.y * a.v.vx + b.p.y * b.v.vx, a.p.x * a.v.vz - b.p.x * b.v.vz - a.p.z * a.v.vx + b.p.z * b.v.vx}
}

func main() {
    parseData, _ := os.ReadFile("../24.1/input.24.1.txt")
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
        x, _ := strconv.ParseInt(strings.TrimSpace(strings.Split(pos, ",")[0]), 10, 64)
        y, _ := strconv.ParseInt(strings.TrimSpace(strings.Split(pos, ",")[1]), 10, 64)
        z, _ := strconv.ParseInt(strings.TrimSpace(strings.Split(pos, ",")[2]), 10, 64)
        vx, _ := strconv.ParseInt(strings.TrimSpace(strings.Split(vel, ",")[0]), 10, 64)
        vy, _ := strconv.ParseInt(strings.TrimSpace(strings.Split(vel, ",")[1]), 10, 64)
        vz, _ := strconv.ParseInt(strings.TrimSpace(strings.Split(vel, ",")[2]), 10, 64)
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
 
    var m []float64
    var bma []float64
    for i := 0; i < len(hailarr); i++ {
        for j := i; j < len(hailarr); j++ {
            cf, b := coeff(hailarr[i], hailarr[j])
            for _, v := range cf[0] {
                m = append(m, float64(v))
            }
            for _, v := range cf[1] {
                m = append(m, float64(v))
            }
            bma = append(bma, float64(b[0]))
            bma = append(bma, float64(b[1]))
        }
    }

    //fmt.Println(len(m), len(bma))
    A := mat.NewDense(len(bma), 6, m)
    B := mat.NewVecDense(len(bma), bma)
   
    var x mat.VecDense
    if err := x.SolveVec(A, B); err != nil {
        fmt.Println(err)
        return
    }
    //fmt.Println(x.RawVector().Data)

    c := 0
    ans := 0
    for _, v := range x.RawVector().Data {
        if c > 2 {
            break
        }
        fmt.Printf("%f\n", v)
        ans += int(v)
        c++
    }
    fmt.Println(ans)
}
