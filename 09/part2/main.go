package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Data struct {
        x float64
        y float64
}
type Edge struct {
        a Data
        b Data        
}

// point p is inside polygon formed by edges e
func is_inside(e []Edge, p Data) bool {
    cnt := 0
    for _, edge := range e {
        a := edge.a
        b := edge.b
        
        if ( ((p.y < a.y) != (p.y < b.y)) && ( p.x < (a.x + (p.y - a.y) * (b.x - a.x) / (b.y - a.y)) ) ) {
            cnt++
        }
    }
    return cnt%2 == 1
}

func ccw(a, b, c Data) bool {
    return (c.y - a.y) * (b.x - a.x) > (b.y - a.y) * (c.x - a.x)
}
// line intersection check
func is_intersect(e []Edge, c Data, d Data) bool {
    for _, edge := range e {
        a := edge.a
        b := edge.b
        if (ccw(a, c, d) != ccw(b, c, d)) && (ccw(a, b, c) != ccw(a, b, d)) {
            return true
        }
        
    }
    return false
}


func main() {
	file, err := os.Open("data/inp.txt")
    // file, err := os.Open("data/check1.txt") // 40
    // file, err := os.Open("data/ex.txt")
    if err != nil {
        panic(err)
    }

    scanner := bufio.NewScanner(file)


    lists := []Data{}
    for scanner.Scan() {
        line := scanner.Text()
		// fmt.Println(line)

        strs := strings.Split(line, ",")
        x, _ := strconv.ParseFloat(strs[0], 64)
        y, _ := strconv.ParseFloat(strs[1], 64)
        lists = append(lists, Data{x: x, y: y})

    }

    edges := []Edge{}
    size := len(lists)
    ans := 0
    for i:=0; i<size; i++ {
        // should have common x or y
        cur := lists[i]
        next := lists[(i+1)%size]
        if cur.x != next.x && cur.y != next.y {
            fmt.Println("error in input")
            return
        }
        // ans = max(ans, int( (next.x - cur.x + 1) * (next.y - cur.y+1) ))
        edges = append(edges, Edge{a: lists[i], b: lists[(i+1)%size]})
    }

    // fmt.Println("edges:", edges)

    // fmt.Println("checking points...", is_inside(edges, Data{x: 11 - 0.5, y: 1 + 0.5}))
    
    for i:=0; i<size; i++ {
        for j:=i+1; j<size; j++ {
            
            // fmt.Println("checking", lists[i], lists[j])
            min_x := min(lists[i].x, lists[j].x)
            max_x := max(lists[i].x, lists[j].x)
            min_y := min(lists[i].y, lists[j].y)
            max_y := max(lists[i].y, lists[j].y)
            offset := 0.5
            
            // check all four corners
            p1 := Data{x: min_x + offset, y: min_y + offset}
            p2 := Data{x: min_x + offset, y: max_y - offset}
            p3 := Data{x: max_x - offset, y: min_y + offset}
            p4 := Data{x: max_x - offset, y: max_y - offset}

            if !is_inside(edges, p1) {
                // fmt.Println("corner1 not inside", p1)
                continue
            }
            if !is_inside(edges, p2) {
                // fmt.Println("corner2 not inside", p2)
                continue
            }
            if !is_inside(edges, p3) {
                // fmt.Println("corner3 not inside", p3)
                continue
            }
            if !is_inside(edges, p4) {
                // fmt.Println("corner4 not inside", p4)
                continue
            }
            
             if is_intersect(edges, p1, p2) {
                continue
            }

            if is_intersect(edges, p3, p4) {
                continue
            }
            if is_intersect(edges, p1, p3) {
                continue
            }
            if is_intersect(edges, p2, p4) {
                continue
            }

            dx := max_x - min_x + 1
            dy := max_y - min_y + 1
            // fmt.Println("found:", min_x, min_y, max_x, max_y, "area:", int(dx*dy))


            ans = max(ans, int(dx*dy))

        }
    }

    fmt.Println(ans)
}
