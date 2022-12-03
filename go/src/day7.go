package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
    "strconv"
    "math"
)

func main() {
    file, err := os.Open("input.txt")

    if err != nil {
        fmt.Println("Error reading the file")
        return
    }
    defer file.Close()

    sc := bufio.NewScanner(file)

    sc.Scan()
    inputLine := strings.Split(sc.Text(), ",")
    crabPositions := make([]int, 0)

    for i := 0; i < len(inputLine); i++ {
        crabPosition, _ := strconv.Atoi(string(inputLine[i]))
        crabPositions = append(crabPositions, crabPosition)
    }

    min, max := crabPositions[0], crabPositions[0]

    for _, position := range crabPositions {

        if position > max {
            max = position
        }

        if position < min {
            min = position
        }
    }

    fuelCost := calculateFuelCost(crabPositions, min, max)

    fmt.Printf("The minimum cost of fuel is %d\n", fuelCost)
}

func calculateFuelCost(positions []int, min, max int) int {
    fuelCosts := make([]int, 0)
    for i := min; i <= max; i++ {
        cost := 0

        for _, position := range positions {
            cost += summation(int(math.Abs(float64(position - i))))
        }
        fuelCosts = append(fuelCosts, cost)
    }

    return minimumFuelCost(fuelCosts)
}

func minimumFuelCost(fuelCosts []int) int {
    min := fuelCosts[0]

    for _, cost := range fuelCosts {
        if cost < min {
            min = cost
        }
    }

    return min
}

func summation(number int) int {
    result := 0
    for i := 0; i <= number; i++ {
        result += i
    }

    return result
}
