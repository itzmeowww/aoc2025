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

    summ := []int{}
    for j := 0; j < len(grid[0]); j++ {
        if grid[len(grid)-1][j] == '|' {
            summ = append(summ, 1)
        } else {
            summ = append(summ, 0)
        }
    }


    for i := len(grid)-1; i >= 0; i-- {
        new_summ := make([]int, len(summ))
        for j := 0; j < len(grid[i]); j++ {
            if grid[i][j] == 'S' {
                fmt.Println(summ[j])
            } else if grid[i][j] == '|' && i+1 < len(grid) && grid[i+1][j] == '^' {
                if j > 0 {
                new_summ[j] += summ[j+1]
                }
                if j < len(grid[i])-1 {
                new_summ[j] += summ[j-1]
                }
            } else if grid[i][j] == '|' {
                new_summ[j] = summ[j]
            }

        } 
        
        summ = new_summ    
    }
}
