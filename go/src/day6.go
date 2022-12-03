package main

import (
    "fmt"
    "bufio"
    "os"
)

func main() {
    file, err := os.Open("input.txt")

    if err != nil {
        fmt.Println("Error reading the file")
        return
    }
    defer file.Close()

    sc := bufio.NewScanner(file)

    lanternfishesPerDay := make(map[byte]int, 0)

    sc.Scan()
    for i := 0; i < len(sc.Text()); i++ {
        if sc.Text()[i] != ',' {
            lanternfishesPerDay[sc.Text()[i]]++
        }
    }


    for i := 0; i < 256; i++ {
        newbornLanternfishes := lanternfishesPerDay['0']

        for j := '0'; j <= '8'; j++ {
            lanternfishesPerDay[byte(j)] = lanternfishesPerDay[byte(j+1)]
        }

        lanternfishesPerDay['6'] += newbornLanternfishes
        lanternfishesPerDay['8'] += newbornLanternfishes
    }

    totalNumberOfLanternfishes := 0

    for _, totalNumberPerDay := range lanternfishesPerDay {
        totalNumberOfLanternfishes += totalNumberPerDay
    }

    fmt.Printf("The total number of lanternfishes is %d\n", totalNumberOfLanternfishes)
}
