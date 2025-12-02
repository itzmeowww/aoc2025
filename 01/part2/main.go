package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
    // file, err := os.Open("data/ex.txt")
	// file, err := os.Open("data/ex2.txt")
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
				if (cur != 0) && (amount%100 >= cur) {
					ans++
				}
				cur -= amount
		case 'R':		
				if (cur != 0) && (amount%100 >= (100 - cur)) {
					ans++
				}
				cur += amount
		}

		ans+= (amount)/100
		cur %= 100
		cur += 100
		cur %= 100

	
		// fmt.Println(cur, ans)
    }
	fmt.Println(ans)
}
