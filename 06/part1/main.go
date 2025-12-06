package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// file, err := os.Open("data/ex.txt")
    // num_lines := 3
    file, err := os.Open("data/inp.txt")
    num_lines := 4
    if err != nil {
        panic(err)
    }

    scanner := bufio.NewScanner(file)
    
    data := [][]int{}

    for i := 0; i < num_lines; i++ {
        scanner.Scan()
        line := scanner.Text()
        data = append(data, []int{})
        for _, part := range strings.Fields(line) {
            num, _ := strconv.Atoi(part)
            data[i] = append(data[i], num)
            // fmt.Println(num)
        }
    }

    fmt.Println(data)

    op := []string{}
    scanner.Scan()
    line := scanner.Text()
    for _, part := range strings.Fields(line) {
        op = append(op, part)
    }
    ans := 0
    for i := 0; i < len(op); i++ {
        switch op[i] {
            case "+":
                res := 0
                for j := 0; j < num_lines; j++ {
                    res += data[j][i]
                }
                ans += res 
            case "*":
                res := 1
                for j := 0; j < num_lines; j++ {
                    res *= data[j][i]
                }
                ans += res 
        }

    }
    fmt.Println(ans)
}
