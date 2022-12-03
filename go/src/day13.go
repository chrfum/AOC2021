package main

import (
    "fmt"
    "os"
    "bufio"
    "strconv"
    "strings"
)

type coordinate struct {
    x, y int
}

func main() {
    file, err := os.Open("input.txt")

    if err != nil {
        fmt.Fprintln(os.Stderr, "Error reading the file")
        return
    }
    defer file.Close()

    sc := bufio.NewScanner(file)

    coordinates := make([]coordinate, 0)
    totalFolds := make([]string, 0)

    var maxX, maxY int

    for sc.Scan() {
        inputRow := sc.Text()

        if len(inputRow) == 0 {
            continue
        } else if len(inputRow) >= 14 {
            totalFolds = append(totalFolds, inputRow)
        } else {
            dotCoordinates := strings.Split(inputRow, ",")

            var coordToAdd coordinate

            coordToAdd.x, _ = strconv.Atoi(dotCoordinates[0])
            coordToAdd.y, _ = strconv.Atoi(dotCoordinates[1])

            if maxX < coordToAdd.x {
                maxX = coordToAdd.x
            }

            if maxY < coordToAdd.y {
                maxY = coordToAdd.y
            }
            coordinates = append(coordinates, coordToAdd)
        }

    }

    coordMap := make(map[coordinate]bool, 0)

    for _, singleCoord := range coordinates {
        coordMap[singleCoord] = true
    }

    for _, singleFold := range totalFolds {

        if string(singleFold[11]) == "x" {
            maxX = maxX / 2

            for key := range coordMap {
                if key.x > maxX {
                    delete(coordMap, key)
                    var temporaryCoordinate coordinate
                    temporaryCoordinate.x = maxX - (key.x - maxX)

                    temporaryCoordinate.y = key.y

                    coordMap[temporaryCoordinate] = true
                }
            }
        } else {
            maxY = maxY / 2

            for key := range coordMap {
                if key.y > maxY {
                    delete(coordMap, key)
                    var temporaryCoordinate coordinate
                    temporaryCoordinate.y = maxY - (key.y - maxY)

                    temporaryCoordinate.x = key.x

                    coordMap[temporaryCoordinate] = true
                }
            }
        }

    }

    dotCounter := len(coordMap)

    fmt.Printf("The total number of dots is %d\n", dotCounter)

    for i := 0; i <= maxY; i++ {
        for j := 0; j <= maxX; j++ {
            var point coordinate = coordinate{j, i}

            if coordMap[point] {
                fmt.Print("#")
            } else {
                fmt.Print(".")
            }
        }

        fmt.Println()
    }

}
