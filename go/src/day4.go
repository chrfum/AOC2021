package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
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

    sc.Scan()
    numbersDrawn := strings.Split(sc.Text(), ",")

    boards := make([][][]string, 100)
    boardNumber := -1

    for sc.Scan() {
        if sc.Text() == "" {
            boardNumber++
        } else {
            line := strings.Split(sc.Text(), " ")
            inputLine := make([]string, 0)

            for i := 0; i < len(line); i++ {
                if line[i] != "" {
                    inputLine = append(inputLine, line[i])
                }
            }

            boards[boardNumber] = append(boards[boardNumber], inputLine)
        }
    }

    win := false
    result := 0
    winningBoards := make(map[int]bool, 0)


    for _, numberDrawn := range numbersDrawn {
        for i, currentBoard := range boards {
            for j, currentRow := range currentBoard {
                for k, number  := range currentRow {

                    if numberDrawn == number {
                        boards[i][j][k] = "-1"
                        win = isWinningBoard(boards[i])

                        if win {
                            winningBoards[i] = true
                            win = false
                        }

                        if len(winningBoards) == len(boards) {
                            win = true
                            result = sumUnmarkedNumbers(currentBoard)
                            winningNumber, _ := strconv.Atoi(numberDrawn)
                            result *= winningNumber
                            break
                        }

                    }
                    
                }
                if win { break }
            }
            if win { break }
        }
        if win { break }
    }

    fmt.Printf("The result is %d\n", result)
}

func isWinningBoard(board [][]string) bool {
    win := false
    for i := 0; i < len(board); i++ {
        for j := 0; j < len(board[i]); j++ {

            if board[i][j] == "-1" {
                win = true
            } else {
                win = false
                break
            }

        }
        if win { return win }
    }

    for i := 0; i < len(board); i++ {
        for j := 0; j < len(board[i]); j++ {

            if board[j][i] == "-1" {
                win = true
            } else {
                win = false
                break
            }

        }
        if win { return win }
    }

    return false
}

func sumUnmarkedNumbers(board [][]string) int {
    sum := 0

    for i := 0; i < len(board); i++ {
        for j := 0; j < len(board[i]); j++ {

            if board[i][j] != "-1" {
                unmarkedNumber, _ := strconv.Atoi(board[i][j])
                sum += unmarkedNumber
            }

        }
    }

    return sum
}
