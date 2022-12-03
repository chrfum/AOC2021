package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
    "unicode"
)

func main() {
    file, err := os.Open("input.txt")

    if err != nil {
        fmt.Fprintln(os.Stderr, "Error reading the file")
        return
    }
    defer file.Close()

    sc := bufio.NewScanner(file)
    pathMap := make(map[string][]string, 0)

    for sc.Scan() {
        singlePath := strings.Split(sc.Text(), "-")

        pathMap[singlePath[0]] = append(pathMap[singlePath[0]], singlePath[1])
        pathMap[singlePath[1]] = append(pathMap[singlePath[1]], singlePath[0])
    }

    fmt.Printf("The total number of paths is %d\n", findPath(pathMap, "start", 0, ""))
}

func findPath(pathMap map[string][]string, cave string, counter int, path string) int {
    if passedTwoTimes(path) {
        return counter
    }

    if cave == "end" {
        counter++
        return counter
    }

    path += cave + "-"
    for i := 0; i < len(pathMap[cave]); i++ {

        if pathMap[cave][i] != "start" {

            if len(pathMap) == 1 && strings.Contains(path, pathMap[cave][i]) && unicode.IsLower(rune(pathMap[cave][i][0])) {
                if strings.Count(path, pathMap[cave][i]) == 2 {
                    return counter
                } else {
                    counter = findPath(pathMap, pathMap[cave][i], counter, path)
                }

            } else  {

                if strings.Contains(path, pathMap[cave][i]) && unicode.IsUpper(rune(pathMap[cave][i][0])) {
                    counter = findPath(pathMap, pathMap[cave][i], counter, path)
                } else {
                    if strings.Count(path, pathMap[cave][i]) == 2 && unicode.IsLower(rune(pathMap[cave][i][0])) {
                        if passedTwoTimes(path) {
                            return counter
                        }
                    } else {
                        counter = findPath(pathMap, pathMap[cave][i], counter, path)
                    }
                }
            }
        }
    }

    return counter
}

func passedTwoTimes(path string) bool {
    caves := strings.Split(path, "-")
    cavesCounter := make(map[string]int, 0)
    counter := 0

    for i := 0; i < len(caves); i++ {
        cavesCounter[caves[i]]++
        if cavesCounter[caves[i]] == 2 && unicode.IsLower(rune(caves[i][0])) {
            counter++
        }
        if counter >= 2 {
            return true
        }
    }


    return false
}
