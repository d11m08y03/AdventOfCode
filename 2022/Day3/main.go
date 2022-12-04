package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)
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

func calculatePoints(chars string) int {
    points := 0
    for _, val := range chars {
        if val <= 122 && val >= 97 {
            points += 1 + int(val) - 97
        } else {
            points += 27 + int(val) - 65
        }
    }

    return points
}

func partOne(fileContents []string) int {

    uniqueChars := ""
    alreadyPassedThrough := ""

    for i := 0; i < len(fileContents); i++ {
        firstHalf := fileContents[i][0 : len(fileContents[i]) / 2]
        secondHalf := fileContents[i][len(fileContents[i]) / 2 : len(fileContents[i])]

        for j := 0; j < len(firstHalf); j++ {
            charFromFirstHalf := string(firstHalf[j])

            for k := 0; k < len(secondHalf); k++ {
                if charFromFirstHalf == string(secondHalf[k]) && !strings.Contains(alreadyPassedThrough, charFromFirstHalf) {
                    uniqueChars += charFromFirstHalf
                    alreadyPassedThrough += charFromFirstHalf
                }
            } 
        }

        firstHalf = ""
        secondHalf = ""
        alreadyPassedThrough = ""
    }

    return calculatePoints(uniqueChars)
}

func partTwo(fileContents []string) int {
    charsFound := ""
    for i := 0; i < len(fileContents); i += 3 {
        firstLine := fileContents[i]
        secondLine := fileContents[i + 1]
        thirdLine := fileContents[i + 2]

        for j := 0; j < len(firstLine); j++ {
            charFromFirstLine := string(firstLine[j])
            if strings.Contains(secondLine, charFromFirstLine) && strings.Contains(thirdLine, charFromFirstLine) {
                charsFound += charFromFirstLine
                break
            }
        }
    }

    return calculatePoints(charsFound)
}

func main() {
    fileContents := readFile("input.txt")
    fmt.Println(partOne(fileContents))
    fmt.Println(partTwo(fileContents))
}
