package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// number of digits in an integer
func dig_len(n int) int {
    ret := 0
    for n > 0 {
        n = n / 10
        ret++
    }
    return ret
}

func main() {
	file, err := os.Open("data/inp.txt")
    // file, err := os.Open("data/ex.txt")
    if err != nil {
        panic(err)
    }

    scanner := bufio.NewScanner(file)
    scanner.Scan()
    line := scanner.Text()
    ranges := strings.Split(line, ",")
    
    sum := 0
    for _, r := range ranges {
        parts := strings.Split(r, "-")
        start, _ := strconv.Atoi(parts[0])
        end, _ := strconv.Atoi(parts[1])
        // fmt.Printf("Range: %d - %d\n", start, end)
        for i := dig_len(start); i <= dig_len(end); i++ {
            if i % 2 == 1 {
                continue
            }
            minn := int(math.Max(float64(start), math.Pow10(i-1)))
            maxx := int(math.Min(float64(end), math.Pow10(i)-1))
            // fmt.Printf("From %d to %d\n", int(minn), int(maxx))

            for j := (minn / int(math.Pow10(i / 2))); j <= int(maxx / int(math.Pow10(i / 2))); j++ {
                num := j * int(math.Pow10(i / 2)) + j;
                if num >= minn && num <= maxx {
                    // fmt.Println(num)
                    sum += num;
                }
            }
        }
        
        
    }
    fmt.Println(sum)
}
