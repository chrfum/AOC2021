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
        fmt.Println("Error reading the file")
        return
    }
    defer file.Close()

    sc := bufio.NewScanner(file)

    totalMeasurments := make([]int, 0)

    for sc.Scan() {
        depthMeasurment, _ := strconv.Atoi(sc.Text())
        totalMeasurments = append(totalMeasurments, depthMeasurment)
    }

    countOfIncreases := 0

    for i := 1; i < len(totalMeasurments) - 2; i++ {
        prev := i - 1
        if sumOfThreeMeasurment(totalMeasurments[i:i+3]) > sumOfThreeMeasurment(totalMeasurments[prev:prev+3]) {
            countOfIncreases++
        }
    }

    fmt.Printf("Totale increases: %d\n", countOfIncreases)
}

func sumOfThreeMeasurment(threeMeasurments []int) int {
    sum := 0
    for i := 0; i < len(threeMeasurments); i++ {
        sum += threeMeasurments[i]
    }
    return sum
}
