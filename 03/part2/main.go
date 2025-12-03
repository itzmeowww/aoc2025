package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func solve(input string) int {
    // fmt.Println(input)
    var arr [13]int
    // arr[j] = max value with j digits
    maxx := 0
    for i := len(input)-1; i>=0; i-- {
        
        cur, _ := strconv.Atoi(string(input[i]))
        // fmt.Println(arr)

        for j := 11; j >= 0; j-- {
            if j >= len(input) - i {
                continue
            }

            arr[j+1] = int(math.Max(float64(arr[j+1]), float64(arr[j] + cur * int(math.Pow10(j)))))
            if j == 11 && arr[j+1] > maxx {
                maxx = arr[j+1]
            }
        }
    }
    return maxx
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
        // fmt.Println(res)
        ans += res
    }

    fmt.Println(ans)
}
