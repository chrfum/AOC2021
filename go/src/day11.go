package main

import (
    "fmt"
    "os"
    "bufio"
    "strconv"
)

func main() {
    file, err := os.Open("input.txt")

    if err != nil {
        fmt.Fprintln(os.Stderr, "Error reading the file")
        return
    }
    defer file.Close()

    sc := bufio.NewScanner(file)
    octopusMap := make([][]int, 0)
    firstEntry := true
    borderLine := make([]int, 0)

    for sc.Scan() {
        if firstEntry {
            for i := 0; i < len(sc.Text()) + 2; i++ {
                borderLine = append(borderLine, 11)
            }
            octopusMap = append(octopusMap, borderLine)
            firstEntry = false
        }

        octopusLine := make([]int, 0)
        octopusLine = append(octopusLine, 11)

        for i := 0; i < len(sc.Text()); i++ {
            energyLevel, _ := strconv.Atoi(string(sc.Text()[i]))
            octopusLine = append(octopusLine, energyLevel)
        }
        octopusLine = append(octopusLine, 11)

        octopusMap = append(octopusMap, octopusLine)
    }

    octopusMap = append(octopusMap, borderLine)
    flashes := 0
    stepCounter := 0

    for !(synchronizedFlash(octopusMap)) {
        stepCounter++
        for j := 1; j < len(octopusMap) - 1; j++ {
            for k := 1; k < len(octopusMap[j]) - 1; k++ {
                octopusMap[j][k]++
                if octopusMap[j][k] == 10 {
                    octopusMap = findFlashes(octopusMap, j, k)
                }
            }
        }

        for j := 1; j < len(octopusMap) - 1; j++ {
            for k := 1; k < len(octopusMap[j]) - 1; k++ {
                if octopusMap[j][k] >= 10 {
                    octopusMap[j][k] = 0
                    flashes++
                }
            }
        }

        if stepCounter == 100 {
            fmt.Printf("The total number of flashes at step 100 is %d\n", flashes)
        }
    }

    fmt.Printf("The first step during which all octopuses flash is %d\n", stepCounter)
}

func findFlashes(matrix [][]int, j, k int) [][]int {
    if matrix[j-1][k-1] < 10 {
        matrix[j-1][k-1]++
        if matrix[j-1][k-1] == 10 {
            matrix = findFlashes(matrix, j-1, k-1)
        }
    }

    if matrix[j][k-1] < 10 {
        matrix[j][k-1]++
        if matrix[j][k-1] == 10 {
            matrix = findFlashes(matrix, j, k-1)
        }
    }

    if matrix[j+1][k-1] < 10 {
        matrix[j+1][k-1]++
        if matrix[j+1][k-1] == 10 {
            matrix = findFlashes(matrix, j+1, k-1)
        }
    }

    if matrix[j-1][k] < 10 {
        matrix[j-1][k]++
        if matrix[j-1][k] == 10 {
            matrix = findFlashes(matrix, j-1, k)
        }
    }

    if matrix[j+1][k] < 10 {
        matrix[j+1][k]++
        if matrix[j+1][k] == 10 {
            matrix = findFlashes(matrix, j+1, k)
        }
    }

    if matrix[j-1][k+1] < 10 {
        matrix[j-1][k+1]++
        if matrix[j-1][k+1] == 10 {
            matrix = findFlashes(matrix, j-1, k+1)
        }
    }

    if matrix[j][k+1] < 10 {
        matrix[j][k+1]++
        if matrix[j][k+1] == 10 {
            matrix = findFlashes(matrix, j, k+1)
        }
    }

    if matrix[j+1][k+1] < 10 {
        matrix[j+1][k+1]++
        if matrix[j+1][k+1] == 10 {
            matrix = findFlashes(matrix, j+1, k+1)
        }
    }

    return matrix
}

func synchronizedFlash(matrix [][]int) bool {
    for i := 1; i < len(matrix) - 1; i++ {
        for j := 1; j < len(matrix[i]) - 1; j++ {
            if matrix[i][j] != 0 {
                return false
            }
        }
    }
    return true
}
