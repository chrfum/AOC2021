package main

import (
    "fmt"
    "os"
    "bufio"
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
    gammaOrEpsilon := make(map[int]int, 0)
    totalNumbers := make([]string, 0)

    var numberLength int

    for sc.Scan() {
        binaryNumber := sc.Text()

        totalNumbers = append(totalNumbers, binaryNumber)
        numberLength = len(binaryNumber)
        for i := 0; i < numberLength; i++ {
            if binaryNumber[i] == '1' {
                gammaOrEpsilon[i]++
            } else {
                gammaOrEpsilon[i]--
            }
        }
    }

    gammaRate, epsilonRate := 0, 0

    for i := 0; i < numberLength; i++ {
        bitPosition := float64(numberLength - i - 1)
        if gammaOrEpsilon[i] > 0 {
            gammaRate += int(math.Pow(2, bitPosition))
        } else {
            epsilonRate += int(math.Pow(2, bitPosition))
        }
    }

    fmt.Printf("The power consumption is %d\n", gammaRate * epsilonRate)

    oxygenRate := oxygenGeneratorRating(totalNumbers, 0, numberLength)
    co2Rate := co2ScrubberRating(totalNumbers, 0, numberLength)
    
    fmt.Printf("The life support rating is %d\n", oxygenRate * co2Rate)
}

func binToDecimal(binary string, numberLength int) int {
    decimalNumber := 0

    for i := 0; i < numberLength; i++ {
        binaryDigit := binary[i]
        bitPosition := float64(numberLength - i - 1)
        if binaryDigit == '1' {
            decimalNumber += int(math.Pow(2, bitPosition))
        }

    }

    return decimalNumber
}

func oxygenGeneratorRating(binaryNumbers []string, digit, maxLength int) int {
    if len(binaryNumbers) == 1 {
        return binToDecimal(binaryNumbers[0], maxLength)
    }

    mostCommonDigit := 0

    for i := 0; i < len(binaryNumbers); i++ {
        if binaryNumbers[i][digit] == '1' {
            mostCommonDigit++
        } else {
            mostCommonDigit--
        }
    }

    binaryNumbersFiltered := make([]string, 0)
    if mostCommonDigit >= 0 {
        for i := 0; i < len(binaryNumbers); i++ {
            if binaryNumbers[i][digit] == '1' {
                binaryNumbersFiltered = append(binaryNumbersFiltered, binaryNumbers[i])
            }
        }
    } else {
        for i := 0; i < len(binaryNumbers); i++ {
            if binaryNumbers[i][digit] == '0' {
                binaryNumbersFiltered = append(binaryNumbersFiltered, binaryNumbers[i])
            }
        }
    }

    digit++
    if digit == maxLength {
        digit = 0
    }
    return oxygenGeneratorRating(binaryNumbersFiltered, digit, maxLength)
}

func co2ScrubberRating(binaryNumbers []string, digit, maxLength int) int {
    if len(binaryNumbers) == 1 {
        return binToDecimal(binaryNumbers[0], maxLength)
    }

    mostCommonDigit := 0

    for i := 0; i < len(binaryNumbers); i++ {
        if binaryNumbers[i][digit] == '1' {
            mostCommonDigit++
        } else {
            mostCommonDigit--
        }
    }

    binaryNumbersFiltered := make([]string, 0)
    if mostCommonDigit < 0 {
        for i := 0; i < len(binaryNumbers); i++ {
            if binaryNumbers[i][digit] == '1' {
                binaryNumbersFiltered = append(binaryNumbersFiltered, binaryNumbers[i])
            }
        }
    } else {
        for i := 0; i < len(binaryNumbers); i++ {
            if binaryNumbers[i][digit] == '0' {
                binaryNumbersFiltered = append(binaryNumbersFiltered, binaryNumbers[i])
            }
        }
    }

    digit++
    if digit == maxLength {
        digit = 0
    }
    return co2ScrubberRating(binaryNumbersFiltered, digit, maxLength)
}
