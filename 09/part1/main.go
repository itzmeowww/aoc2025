package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Data struct {
        x int
        y int
}

func main() {
	file, err := os.Open("data/inp.txt")
    // file, err := os.Open("data/ex.txt")
    if err != nil {
        panic(err)
    }

    scanner := bufio.NewScanner(file)
    
    lists := []Data{}
    for scanner.Scan() {
        line := scanner.Text()
		// fmt.Println(line)

        strs := strings.Split(line, ",")
        x, _ := strconv.Atoi(strs[0])
        y, _ := strconv.Atoi(strs[1])
        lists = append(lists, Data{x: x, y: y})
    }

    ans:= 0
    size := len(lists)

    for i:=0; i<size; i++ {
        for j:=i+1; j<size; j++ {
            dx := math.Abs(float64(lists[i].x - lists[j].x))+1
            dy := math.Abs(float64(lists[i].y - lists[j].y))+1

            ans = max(ans, int(dx*dy))
        }
    }

    fmt.Println(ans)
    
}
