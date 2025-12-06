package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
    
    data := []string{}

    for i := 0; i < num_lines+1; i++ {
        scanner.Scan()
        line := scanner.Text()
        data = append(data, line)
    }
    fmt.Println(data)
    idx := 0
    ans := 0
    for true {
        if idx >= len(data[num_lines]) {
            break
        }
        // fmt.Printf("idx: %d\n", idx)
        // fmt.Println(data[num_lines][idx])
        op := data[num_lines][idx]
        var res int
        if op == '+' {
            res = 0
        } else if op == '*' {
            res = 1
        }
        changed := false
        for true {
            changed = false
            num_str := []byte{}
            for i := 0; i < num_lines; i++ {
                if idx < len(data[i]) && data[i][idx] >= '0' && data[i][idx] <= '9' {
                    num_str = append(num_str, data[i][idx])
                    changed = true
                }
            }
            idx++
            if !changed {
                break
            }
            num , _ := strconv.Atoi(string(num_str))
            // fmt.Printf("num: %d\n", num)
            if op == '+' {
                res += num
            } else if op == '*' {
                res *= num
            }
            have := false
            for i := 0; i < num_lines; i++ {
                if idx < len(data[i]) {
                    have = true
                }
            }
            if !have {
                break
            }
            
        }

        ans += res
    }

    fmt.Println(ans)
    
   
}
