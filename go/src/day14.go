package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
)

func main() {
    file, err := os.Open("input.txt")

    if err != nil {
        fmt.Fprintf(os.Stderr, "Error reading the file")
    }
    defer file.Close()

    sc := bufio.NewScanner(file)
    firstScan := true
    polymerPairs := make(map[string]int, 0)
    insertionRules := make(map[string]byte, 0)
    occurrenceCounter := make(map[byte]int, 0)

    for sc.Scan() {
        if firstScan {
            firstScan = false
            polymer := sc.Text()

            for i := 0; i < len(polymer) - 1; i++ {
                polymerPairs[polymer[i:i+2]]++
                occurrenceCounter[polymer[i]]++
            }
            occurrenceCounter[polymer[len(polymer) - 1]]++

        } else if sc.Text() != "" {
            rules := strings.Split(sc.Text(), " -> ")
            insertionRules[rules[0]] = rules[1][0]
        }
    }

    for i := 0; i < 40; i++ {
        temporaryPairs := make(map[string]int, 0)

        for pair, occurrence := range polymerPairs {
            temporaryPairs[pair] = occurrence
        }

        for pair, char := range insertionRules {
            newChars := temporaryPairs[pair]
            occurrenceCounter[char] += newChars
            polymerPairs[string(pair[0]) + string(char)] += newChars
            polymerPairs[string(char) + string(pair[1])] += newChars
            polymerPairs[pair] -= newChars
        }
    }

    min, max := 0, 0
    firstIter := true

    for _, occurrence := range occurrenceCounter {
        if firstIter {
            firstIter = false
            min = occurrence
            max = occurrence
        } else {

            if occurrence > max {
                max = occurrence
            }

            if occurrence < min {
                min = occurrence
            }
        }
    }

    fmt.Printf("The result is %d\n", max - min)
}
