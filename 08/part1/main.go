package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Point struct{
        x int
        y int
        z int
}

type Data struct{
        idx1 int
        idx2 int
        dis int
}

var pa [1005]int
var ans [1005]int

func fp(x int) int {
    if pa[x] != x {
        pa[x] = fp(pa[x])
    }
    return pa[x]
}

func main() {
	file, err := os.Open("data/inp.txt")
    maxC := 1000
    // file, err := os.Open("data/ex.txt")
    // maxC := 10

    if err != nil {
        panic(err)
    }

    scanner := bufio.NewScanner(file)
    
    lists := []Point{}

    for scanner.Scan() {
        line := scanner.Text()
		strs := strings.Split(line, ",")
        x, _ := strconv.Atoi(strs[0])
        y, _ := strconv.Atoi(strs[1])
        z, _ := strconv.Atoi(strs[2])
        lists = append(lists, Point{x: x, y: y, z: z})
    }

    size := len(lists)

    data := []Data{}
    for i:=0; i<size; i++ {
        pa[i] = i
        for j:=i+1; j<size; j++ {
            dx := lists[i].x - lists[j].x
            dy := lists[i].y - lists[j].y
            dz := lists[i].z - lists[j].z
            dist := dx*dx + dy*dy + dz*dz
            data = append(data, Data{idx1: i, idx2: j, dis: dist})
        }
    }

    sort.Slice(data, func(i, j int) bool {
        return data[i].dis < data[j].dis
    })

    cou := 0
    for _, d := range data {
        d1 := fp(d.idx1)
        d2 := fp(d.idx2)
        if d1 != d2 {
            // fmt.Println("connect", d.idx1, d.idx2)
            pa[d1] = d2
            
            
        }
        cou++
        if cou == maxC {
                break
            }
    }

    for i:=0; i<size; i++ {
        // fmt.Println(i, "->", fp(i))
        ans[fp(i)]++
    }

    // for i:=0; i<size; i++ {
    //     fmt.Println(ans[fp(i)])
    // }


    
    sort.Slice(ans[:], func(i, j int) bool {
        return ans[i] > ans[j]
    })


    // fmt.Println("Part 1: ", ans[0])
    // fmt.Println("Part 2: ", ans[1])
    // fmt.Println("Part 3: ", ans[2])

    fmt.Println(ans[0] * ans[1] * ans[2])

}
