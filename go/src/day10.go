package main

import (
    "fmt"
    "os"
    "bufio"
    "sort"
)

func main() {
    file, err := os.Open("input.txt")

    if err != nil {
        fmt.Fprintln(os.Stderr, "Error reading the file")
        return
    }
    defer file.Close()

    sc := bufio.NewScanner(file)
    syntaxErrorScore := 0
    allCompleteScores := make([]int, 0)

    for sc.Scan() {
        line := sc.Text()
        chunkStack := make([]rune, 0)
        legalLine := true

        for _, char := range line {
            if char == '(' || char == '[' || char == '{' || char == '<' {
                chunkStack = append(chunkStack, char)
            } else {
                charToVerify := chunkStack[len(chunkStack)-1]
                if char == ')' && charToVerify == '(' ||
                char == ']' && charToVerify == '[' ||
                char == '}' && charToVerify == '{' ||
                char == '>' && charToVerify == '<' {
                    chunkStack = chunkStack[:len(chunkStack)-1]
                } else {
                    switch char {
                    case ')':
                        syntaxErrorScore += 3
                    case ']':
                        syntaxErrorScore += 57
                    case '}':
                        syntaxErrorScore += 1197
                    case '>':
                        syntaxErrorScore += 25137
                    }
                    legalLine = false
                    break
                }
            }
        }

        if legalLine {
            completeScore := 0
            remainingChars := len(chunkStack)

            for i := 0; i < remainingChars; i++ {

                switch chunkStack[len(chunkStack)-1] {
                case '(':
                    completeScore  = (completeScore * 5) + 1
                case '[':
                    completeScore  = (completeScore * 5) + 2
                case '{':
                    completeScore  = (completeScore * 5) + 3
                case '<':
                    completeScore  = (completeScore * 5) + 4
                }

                chunkStack = chunkStack[:len(chunkStack)-1]
            }

            allCompleteScores = append(allCompleteScores, completeScore)
        }
    }

    sort.Ints(allCompleteScores)
    middleScore := allCompleteScores[len(allCompleteScores)/2]

    fmt.Printf("The syntax error score is %d\n", syntaxErrorScore)
    fmt.Printf("The middle score is %d\n", middleScore)
}
