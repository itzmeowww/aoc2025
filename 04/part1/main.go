package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// file, err := os.Open("data/ex.txt")
    file, err := os.Open("data/inp.txt")
    if err != nil {
        panic(err)
    }

    scanner := bufio.NewScanner(file)
    grid := []string{}
    for scanner.Scan() {
        line := scanner.Text()
		// fmt.Println(line)
        grid = append(grid, line)
    }

    ii := [8]int{-1, -1, -1, 0, 0, 1, 1, 1}
    jj := [8]int{-1, 0, 1, -1, 1, -1, 0, 1}
    maxi := len(grid)
    maxj := len(grid[0])
    ans := 0
    for i := 0; i < maxi; i++ {
        for j := 0; j < maxj; j++ {
            if grid[i][j] != '@' {
                continue
            }
            cou := 0
            for k:= 0; k < 8; k++ {
                ni := i + ii[k]
                nj := j + jj[k]
                if (ni >= 0 && ni < maxi && nj >= 0 && nj < maxj) {
                    if grid[ni][nj] == '@' {
                        cou++
                    }
                }
            }
            if cou < 4 {
                ans++
            }
        }
    }
    fmt.Println(ans)
}
