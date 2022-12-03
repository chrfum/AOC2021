package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
    "strconv"
)

func main() {
    file, err := os.Open("input.txt")

    if err != nil {
        fmt.Println("Error reading the file")
        return
    }
    defer file.Close()

    sc := bufio.NewScanner(file)

    linesPerPoint := make(map[[2]int]int, 0)

    for sc.Scan() {
        points := strings.Split(sc.Text(), " -> ")

        firstPoint := strings.Split(points[0], ",")
        x1, _ := strconv.Atoi(firstPoint[0])
        y1, _ := strconv.Atoi(firstPoint[1])

        secondPoint := strings.Split(points[1], ",")
        x2, _ := strconv.Atoi(secondPoint[0])
        y2, _ := strconv.Atoi(secondPoint[1])

        linesPerPoint = distanceBetweenPoints(x1, y1, x2, y2, linesPerPoint)
    }

    overlappingPoints := 0

    for _, lines := range linesPerPoint {
        if lines > 1 {
            overlappingPoints++
        }
    }

    fmt.Printf("The overlappingPoints are %d\n", overlappingPoints)
}

func distanceBetweenPoints(x1, y1, x2, y2 int, lines map[[2]int]int) map[[2]int]int {
    if x1 == x2 {
        distance := 0

        if y1 > y2 {
            distance = y1 - y2

            for i := 0; i <= distance; i++ {
                var point [2]int = [2]int {x1, y1}
                lines[point]++
                y1--
            }

        } else if y2 > y1 {
            distance = y2 - y1

            for i := 0; i <= distance; i++ {
                var point [2]int = [2]int {x1, y1}
                lines[point]++
                y1++
            }
        }

    } else if y1 == y2 {
        distance := 0

        if x1 > x2 {
            distance = x1 - x2

            for i := 0; i <= distance; i++ {
                var point [2]int = [2]int {x1, y1}
                lines[point]++
                x1--
            }

        } else if x2 > x1 {
            distance = x2 - x1

            for i := 0; i <= distance; i++ {
                var point [2]int = [2]int {x1, y1}
                lines[point]++
                x1++
            }
        }

    } else if x1 > x2 {
        distance := x1 - x2

        if y1 > y2 {
            for i := 0; i <= distance; i++ {
                var point [2]int = [2]int {x1, y1}
                lines[point]++
                x1--
                y1--
            }

        } else if y2 > y1 {
            for i := 0; i <= distance; i++ {
                var point [2]int = [2]int {x1, y1}
                lines[point]++
                x1--
                y1++
            }
        }

    } else if x1 < x2 {
        distance := x2 - x1

        if y1 > y2 {
            for i := 0; i <= distance; i++ {
                var point [2]int = [2]int {x1, y1}
                lines[point]++
                x1++
                y1--
            }
            
        } else if y2 > y1 {
            for i := 0; i <= distance; i++ {
                var point [2]int = [2]int {x1, y1}
                lines[point]++
                x1++
                y1++
            }
        }
    }

    return lines
}
