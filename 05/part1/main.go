package main

import (
	"bufio"
	"fmt"
	"os"
    "strconv"
	"strings"
)

type Data struct {
    start int
    end   int
}

func main() {
	file, err := os.Open("data/inp.txt")
    // file, err := os.Open("data/ex.txt")
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
            start: start,
            end:   end,
        })
    }

    ans := 0

    for scanner.Scan() {
        line := scanner.Text()
		num, _ := strconv.Atoi(line)
        for _, r := range ranges {
            if num >= r.start && num <= r.end {
                ans ++
                break
            }
        }
    }

    fmt.Println(ans)
}
