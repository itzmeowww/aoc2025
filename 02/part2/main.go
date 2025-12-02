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

    mp := make(map[int]bool)

    
    for _, r := range ranges {
        parts := strings.Split(r, "-")
        start, _ := strconv.Atoi(parts[0])
        end, _ := strconv.Atoi(parts[1])
        // fmt.Printf("Range: %d - %d\n", start, end)
        for i := dig_len(start); i <= dig_len(end); i++ {
            minn := int(math.Max(float64(start), math.Pow10(i-1)))
            maxx := int(math.Min(float64(end), math.Pow10(i)-1))
            // fmt.Printf("From %d to %d\n", int(minn), int(maxx))
            // split into k parts
            for k := 2; k <= i; k ++ {
                
                if i % k != 0 {
                    continue
                }

                // size of each part
                si := i / k

                // fmt.Printf("split to k = %d parts\n", k)

                for j := (minn / int(math.Pow10(i - (si)))); j <= int(maxx / int(math.Pow10(i - (si)))); j++ {
                    num := 0
                    for m := 0; m < k; m++ {
                        num = num * int(math.Pow10(si)) + j
                    }
                    
                    if num >= minn && num <= maxx {
                        // fmt.Println(num)
                        if _, exists := mp[num]; !exists {
                            mp[num] = true
                            sum += num;
                        }
                        
                    }
                }
            }   
        }        
    }
    fmt.Println(sum)
}
