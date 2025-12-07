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
    
    grid := [][]rune{}

    for scanner.Scan() {
        line := scanner.Text()
        grid = append(grid, []rune(line))
		// fmt.Println(line)
    }
     total := 0

   
    for i := 0; i < len(grid); i++ {
   
        for j := 0; j < len(grid[i]); j++ {

            fmt.Print(string(grid[i][j]))

            if grid[i][j] == 'S' && i < len(grid)-1 && grid[i+1][j] == '.' {
                grid[i+1][j] = '|'
            } else if grid[i][j] == '|' {
                if i < len(grid)-1 && grid[i+1][j] == '.' {
                    grid[i+1][j] = '|'
                } else if i < len(grid)-1 && grid[i+1][j] == '^' {
                    total++
                    if j > 0 && grid[i+1][j-1] == '.' {
                        grid[i+1][j-1] = '|'
                    }
                    if j < len(grid[i])-1 && grid[i+1][j+1] == '.' {
                        grid[i+1][j+1] = '|'
                    }
                }

            }
        }
        fmt.Println()
        
    }
   
   
    fmt.Println(total)

}
