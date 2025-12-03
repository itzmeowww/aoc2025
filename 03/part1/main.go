package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func solve(input string) int {
    if len(input) == 2 {
        val, _ := strconv.Atoi(input)
        return val
    }
    var maxx int = 0
    for i := 1; i < len(input); i++ {
        val, _ := strconv.Atoi(string(input[0]) + string(input[i]))
        if val > maxx {
            maxx = val
        }
    }

    return int(math.Max(float64(maxx), float64(solve(input[1:]))))
}

func main() {
    // file, err := os.Open("data/ex.txt")
	file, err := os.Open("data/inp.txt")
    if err != nil {
        panic(err)
    }

    scanner := bufio.NewScanner(file)
    
    ans := 0
    for scanner.Scan() {
        line := scanner.Text()
		res := solve(line)
        ans += res
    }

    fmt.Println(ans)
}
