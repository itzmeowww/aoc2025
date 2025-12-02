package main

import (
    "bufio"
    "fmt"
    "os"
	"strconv"
)

func main() {
    // file, err := os.Open("data/ex.txt")
	file, err := os.Open("data/inp.txt")
    if err != nil {
        panic(err)
    }

    scanner := bufio.NewScanner(file)
    var cur = 50
	var ans int
    for scanner.Scan() {
        line := scanner.Text()
		dir := line[0]
		amount, _ := strconv.Atoi(line[1:])
		switch dir {
		case 'L':
				cur -= amount
		case 'R':
				cur += amount
		}
		cur %= 100
		cur += 100
		cur %= 100

		if cur == 0 {
			ans++
		}
    }
	fmt.Println(ans)
}
