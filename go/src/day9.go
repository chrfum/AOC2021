package main

import (
    "fmt"
    "os"
    "bufio"
)

type index struct {
    i, j int
}

func main() {
    file, err := os.Open("input.txt")

    if err != nil {
        fmt.Fprintln(os.Stderr, "Error reading the file")
        return
    }
    defer file.Close()

    sc := bufio.NewScanner(file)
    heightmap := make([][]rune, 0)

    borderRow := make([]rune, 0)
    firstIter := true
    for sc.Scan() {
        inputLine := sc.Text()

        if firstIter {
            for i := 0; i < len(inputLine) + 2; i++ {
                borderRow = append(borderRow, '9')
            }
            heightmap = append(heightmap, borderRow)
            firstIter = false
        }

        inputLine = "9" + inputLine + "9"
        currentRow := make([]rune, 0)
        for _, height := range inputLine {
            currentRow = append(currentRow, height)
        }
        heightmap = append(heightmap, currentRow)
    }

    heightmap = append(heightmap, borderRow)

    allBasins := make([]int, 0)
    indexMap := make(map[index]bool, 0)

    for i := 1; i < len(heightmap) - 1; i++ {
        for j := 1; j < len(heightmap[i]) - 1; j++ {
            if isLowPoint(heightmap, i, j, indexMap) {
                indexMap[index{i, j}] = true
                counter := 1
                counter = findBasin(heightmap, i+1, j, counter, indexMap)
                counter = findBasin(heightmap, i-1, j, counter, indexMap)
                counter = findBasin(heightmap, i, j+1, counter, indexMap)
                counter = findBasin(heightmap, i, j-1, counter, indexMap)

                allBasins = append(allBasins, counter)
            }
        }
    }

    fmt.Printf("The result of the three largest basins multiplied is %d\n", maxThree(allBasins))
}

func isLowPoint(matrix [][]rune, i, j int, indexMap map[index]bool) bool {
    if indexMap[index{i, j}] {
        return false
    }

    if ((matrix[i][j] < matrix[i+1][j]) || indexMap[index{i+1, j}]) &&
        ((matrix[i][j] < matrix[i-1][j]) || indexMap[index{i-1, j}]) &&
        ((matrix[i][j] < matrix[i][j+1]) || indexMap[index{i, j+1}]) &&
        ((matrix[i][j] < matrix[i][j-1]) || indexMap[index{i, j-1}]) {
        return true
    } else {
        return false
    }
}

func findBasin(matrix [][]rune, i, j int, counter int, indexMap map[index]bool) int {
    if matrix[i][j] == '9' || indexMap[index{i,j}] {
        indexMap[index{i, j}] = true
        return counter
    }

    if matrix[i][j] != '9' {
        counter++
        indexMap[index{i, j}] = true
        counter = findBasin(matrix, i+1, j, counter, indexMap)
        counter = findBasin(matrix, i-1, j, counter, indexMap)
        counter = findBasin(matrix, i, j+1, counter, indexMap)
        counter = findBasin(matrix, i, j-1, counter, indexMap)
    }

    return counter
}

func maxThree(values []int) int {
    max1, max2, max3 := 0, 0, 0

    for _,  value := range values {
        if value > max3 {
            if value > max2 {

                if value > max1 {
                    max3 = max2
                    max2 = max1
                    max1 = value
                } else {
                    max3 = max2
                    max2 = value
                }

            } else {
                max3 = value
            }
        }
    }

    return max1 * max2 * max3
}
