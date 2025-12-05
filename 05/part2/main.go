package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Data struct {
    start int64
    end   int64
}

func main() {
	file, err := os.Open("data/inp.txt")
    // file, err := os.Open("data/ex.txt")
    // file, err := os.Open("data/ex2.txt")
    if err != nil {
        panic(err)
    }

    scanner := bufio.NewScanner(file)
    var ranges []Data


    for scanner.Scan() {
        line := scanner.Text()
        if line == "" {
            break
        }
		
        parts := strings.Split(line, "-")
        start, _ := strconv.Atoi(parts[0])
        end, _ := strconv.Atoi(parts[1])
        ranges = append(ranges, Data{
            start: int64(start),
            end:   int64(end),
        })
    }

    
    sort.Slice(ranges, func(i, j int) bool {
        if ranges[i].start == ranges[j].start {
            return ranges[i].end < ranges[j].end
        }
        return ranges[i].start < ranges[j].start
    })

    var ans int64 = 0
    last := int64(0)

    for _, r := range ranges {
        
        // fmt.Printf("[%d, %d], last: %d\n", r.start, r.end, last)
        if r.end >= last {
            ans += r.end - max(r.start, last) + 1
        }
        // fmt.Printf("ans: %d\n", ans)
        last = max(last, r.end+1)

    }

    fmt.Println(ans)
}
