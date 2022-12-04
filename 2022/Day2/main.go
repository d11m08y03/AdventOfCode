package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
For part one
A -> Rock -> X
B -> Paper -> Y
C -> Scissors -> Z

Score per round: shapeScore + outcomeScore
shapeScore:         |   outcomeScore:
- Rock -> 1         |   - Win -> 6
- Paper -> 2        |   - Draw -> 3
- Scissors -> 3     |   - Lose -> 0
*/

func processPointsForPartOne(fileContents []string) int {
    points := 0
    for i := 0; i < len(fileContents); i++ {
        plays := strings.Split(fileContents[i], " ")
        switch plays[1] {
        // I played Rock
        case "X":
            points += 1
            switch plays[0] {
            case "A": points += 3
            case "C": points += 6
            }

        // I played Paper
        case "Y":
            points += 2
            switch plays[0] {
            case "A": points += 6
            case "B": points += 3
            }

        // I played Scissors
        case "Z":
            points += 3
            switch plays[0] {
            case "B": points += 6
            case "C": points += 3
            }
        }
    }

    return points
}

/*
For part two
A -> Rock 
B -> Paper 
C -> Scissors 

X -> I need to lose
Y -> End in a draw
Z -> I need to win

Score per round: shapeScore + outcomeScore
shapeScore:         |   outcomeScore:
- Rock -> 1         |   - Win -> 6
- Paper -> 2        |   - Draw -> 3
- Scissors -> 3     |   - Lose -> 0
*/

func processPointsForPartTwo(fileContents []string) int {
    points := 0

    for i := 0; i < len(fileContents); i++ {
        plays := strings.Split(fileContents[i], " ")
        switch plays[0] {
        // Elf played Rock
        case "A":
            switch plays[1] {
            case "X": points += 3
            case "Y": points += 3 + 1
            case "Z": points += 6 + 2
            }
        // Elf played Paper
        case "B":
            switch plays[1] {
            case "X": points += 1
            case "Y": points += 3 + 2
            case "Z": points += 6 + 3
            }
        // Elf played Scissors
        case "C":
            switch plays[1] {
            case "X": points += 2
            case "Y": points += 3 + 3
            case "Z": points += 6 + 1
            }
        }
    }

    return points
}

func readFile(filename string) []string {
    var fileContents []string

    file, err := os.Open(filename)
    if err != nil {
        println(err.Error())
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        fileContents = append(fileContents, scanner.Text())
    }

    return fileContents
}

func main() {
    fileContents := readFile("input.txt")
    fmt.Println(processPointsForPartOne(fileContents))
    fmt.Println(processPointsForPartTwo(fileContents))
}
