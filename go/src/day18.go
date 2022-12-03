package main

import (
    "fmt"
    "os"
    "bufio"
)

type treeNode struct {
    value int
    sx *treeNode
    dx *treeNode
    father *treeNode
}

type tree struct {
    root *treeNode
}

func main() {
    file, err := os.Open("input.txt")

    if err != nil {
        fmt.Fprintln(os.Stderr, "Error reading the file")
        return
    }
    defer file.Close()

    sc := bufio.NewScanner(file)
    var fishTree tree

    firstFish := true
    snailfishNumbers := make([]string, 0)

    for sc.Scan() {
        snailfishNumber := sc.Text()
        snailfishNumbers = append(snailfishNumbers, snailfishNumber)

        if firstFish {
            firstFish = false
            fishTree.root = newNode(-1)
            readTree(snailfishNumber[1:], fishTree, fishTree.root)
        } else {
            secondRoot := newNode(-1)
            readTree(snailfishNumber[1:], fishTree, secondRoot)
            newRoot := newNode(-1)
            newRoot.sx = fishTree.root
            newRoot.dx = secondRoot
            fishTree.root = newRoot

            for {
                _, exploded := explode(fishTree.root, 0, fishTree, false)
                if exploded {
                    continue
                }
                splitted := split(fishTree.root, false)
                if splitted {
                    continue
                } else {
                    break
                }
            }
        }
    }

    for fishTree.root.sx != nil && fishTree.root.dx != nil {
        magnitude(fishTree.root)
    }

    fmt.Printf("The magnitude of the final sum is %d\n", fishTree.root.value)

    magnitudes := make([]int, 0)
    for i := 0; i < len(snailfishNumbers); i++ {
        for j := 0; j < len(snailfishNumbers); j++ {
            if i != j {
                currentFish := "[" + snailfishNumbers[i] + "," + snailfishNumbers[j] + "]"
                fishTree.root = newNode(-1)
                readTree(currentFish[1:], fishTree, fishTree.root)

                for {
                    _, exploded := explode(fishTree.root, 0, fishTree, false)
                    if exploded {
                        continue
                    }
                    splitted := split(fishTree.root, false)
                    if splitted {
                        continue
                    } else {
                        break
                    }
                }

                for fishTree.root.sx != nil && fishTree.root.dx != nil {
                    magnitude(fishTree.root)
                }

                magnitudes = append(magnitudes, fishTree.root.value)
            }
        }
    }

    max := magnitudes[0]
    for i := 1; i < len(magnitudes); i++ {
        if magnitudes[i] > max {
            max = magnitudes[i]
        }
    }

    fmt.Printf("The largest magnitude of any sum of two different snailfish numbers is %d\n", max)
}

func readTree(fish string, fishTree tree, precNode *treeNode) {
    if len(fish) == 0 {
        return
    }

    if fish[0] == '[' {

        if precNode.sx == nil {
            precNode.sx = newNode(-1)
            precNode.sx.father = precNode
            readTree(fish[1:], fishTree, precNode.sx)
        } else {
            precNode.dx = newNode(-1)
            precNode.dx.father = precNode
            readTree(fish[1:], fishTree, precNode.dx)
        }

    } else if fish[0] == ']' {
        readTree(fish[1:], fishTree, precNode.father)

    } else if fish[0] >= '0' && fish[0] <= '9' {
        if precNode.sx == nil {
            precNode.sx = newNode(int(fish[0] - '0'))
            precNode.sx.father = precNode
            readTree(fish[1:], fishTree, precNode)
        } else {
            precNode.dx = newNode(int(fish[0] - '0'))
            precNode.dx.father = precNode
            readTree(fish[1:], fishTree, precNode)
        }

    } else if fish[0] == ',' {
        readTree(fish[1:], fishTree, precNode)
    }

}

func newNode(val int) *treeNode {
    return &treeNode{val, nil, nil, nil}
}

func printTree(root *treeNode) {
    if root != nil {
        if root.value == -1 {
            fmt.Print("[")
            printTree(root.sx)
            fmt.Print(",")
            printTree(root.dx)
            fmt.Print("]")
        } else {
            fmt.Print(root.value)
        }
    }
}

func treeToSlice(root *treeNode, leaves []*treeNode) []*treeNode {
    if root != nil {

        if root.value == -1 {
            leaves = treeToSlice(root.sx, leaves)
            leaves = treeToSlice(root.dx, leaves)
        } else {
            leaves = append(leaves, root)
        }
    }

    return leaves
}

func explode(root *treeNode, counter int, fishTree tree, exploded bool) (int, bool) {
    if exploded {
        return counter, exploded
    }
    if root != nil {
        if root.value == -1 {
            counter++
            counter, exploded = explode(root.sx, counter, fishTree, exploded)
            if exploded {
                return counter, exploded
            }
            counter, exploded = explode(root.dx, counter, fishTree, exploded)
            if exploded {
                return counter, exploded
            }
            counter--
        }

        if counter >= 4 && root.value == -1 && root.sx.value != -1 && root.dx.value != -1 {
            counter = 0
            exploded = true
            sx := root.sx.value
            dx := root.dx.value
            root.sx = nil
            root.dx = nil
            root.value = 0

            leaves := make([]*treeNode, 0)
            leaves = treeToSlice(fishTree.root, leaves)
            for i := 0; i < len(leaves); i++ {
                if leaves[i] == root && i > 0 {
                    leaves[i-1].value += sx
                }
                if leaves[i] == root && i < len(leaves) - 1 {
                    leaves[i+1].value += dx
                }
            }
        }
    }

    return counter, exploded
}

func split(root *treeNode, splitted bool) bool {
    if splitted {
        return splitted
    }
    if root != nil {
        if root.value == -1 {
            splitted = split(root.sx, splitted)
            if splitted {
                return splitted
            }
            splitted = split(root.dx, splitted)
            if splitted {
                return splitted
            }
        } else {
            if root.value > 9 {
                splitted = true
                valueSx := root.value / 2
                valueDx := root.value - valueSx
                root.value = -1
                root.sx = newNode(valueSx)
                root.dx = newNode(valueDx)
                root.sx.father = root
                root.dx.father = root
            }
        }
    }
    return splitted
}

func magnitude(root *treeNode) {
    if root != nil {
        if root.value == -1 && root.sx.value != -1 && root.dx.value != -1 {
            root.value = root.sx.value * 3 + root.dx.value * 2
            root.sx = nil
            root.dx = nil
        }

        if root.value == -1 {
            magnitude(root.sx)
            magnitude(root.dx)
        }
    }
}
