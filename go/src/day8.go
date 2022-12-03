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
        fmt.Fprintln(os.Stderr, "Error reading the file")
        return
    }
    defer file.Close()

    sc := bufio.NewScanner(file)
    sumOfTheOutputValues := 0

    for sc.Scan() {
        line := strings.Split(sc.Text(), "|")
        signalPatterns := strings.Split(line[0], " ")
        outputValues := strings.Split(line[1], " ")

        signalPatterns = signalPatterns[:len(signalPatterns)-1]
        outputValues = outputValues[1:]

        patternTraductions := make(map[int]string, 0)

        for _, pattern := range signalPatterns {

            if len(pattern) == 2 {
                patternTraductions[1] = pattern
            } else if len(pattern) == 3 {
                patternTraductions[7] = pattern
            } else if len(pattern) == 4 {
                patternTraductions[4] = pattern
            } else if len(pattern) == 7 {
                patternTraductions[8] = pattern
            }
        }

        for _, pattern := range signalPatterns {

            if len(pattern) == 5 {

                counter3 := 0
                for i := 0; i < len(patternTraductions[1]); i++ {
                    counter3 += strings.Count(pattern, string(patternTraductions[1][i]))
                }

                if counter3 == len(patternTraductions[1]) {
                    patternTraductions[3] = pattern
                } else {

                    counter2or5 := 0
                    for i := 0; i < len(patternTraductions[4]); i++ {
                        counter2or5 += strings.Count(pattern, string(patternTraductions[4][i]))
                    }
                    if counter2or5 == 3 {
                        patternTraductions[5] = pattern
                    } else {
                        patternTraductions[2] = pattern
                    }
                }
            }
        }

        for _, pattern := range signalPatterns {

            if len(pattern) == 6 {

                counter6 := 0
                for i := 0; i < len(patternTraductions[7]); i++ {
                    counter6 += strings.Count(pattern, string(patternTraductions[7][i]))
                }

                counter3 := 0
                for i := 0; i < len(patternTraductions[3]); i++ {
                    counter3 += strings.Count(pattern, string(patternTraductions[3][i]))
                }


                if counter6 != len(patternTraductions[7]) {
                    patternTraductions[6] = pattern
                } else if counter3 == len(patternTraductions[3]) {
                    patternTraductions[9] = pattern
                } else {
                    patternTraductions[0] = pattern
                }
            }

        }

        finalOutput := 0
        for _, outputValue := range outputValues {
            for traduction, pattern := range patternTraductions {
                if len(outputValue) == len(pattern) {
                    counter := 0
                    for i := 0; i < len(pattern); i++ {
                        counter += strings.Count(outputValue, string(pattern[i]))
                    }

                    if counter == len(outputValue) {
                        finalOutput = (finalOutput * 10) + traduction
                    }
                }
            }
        }

        sumOfTheOutputValues += finalOutput
    }

    fmt.Printf("The sum of the output values is %d\n", sumOfTheOutputValues)
}
