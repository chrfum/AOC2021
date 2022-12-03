package main

import (
    "fmt"
    "os"
    "bufio"
    "strconv"
    "strings"
)

func main() {
    file, err := os.Open("input.txt")

    if err != nil {
        fmt.Println("Error reading the file")
        return
    }
    defer file.Close()

    sc := bufio.NewScanner(file)

    var horizontalPosition, depth, aim int

    for sc.Scan() {
        commandAndValue := strings.Split(sc.Text(), " ")
        command := commandAndValue[0]
        value, _ := strconv.Atoi(commandAndValue[1])

        switch command {
        case "forward":
            horizontalPosition += value
            depth += (aim * value)
        case "down":
            aim += value
        case "up":
            aim -= value
        }
    }

    fmt.Printf("Horizontal position * depth = %d\n", horizontalPosition * depth)
}
