package main

import (
	"bufio"
	"fmt"
	"os"
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

func partOne(fileContents []string) int {
    var initialItems[] byte 
        
    for i := 0; i < len(fileContents); i++ {
        firstHalf := fileContents[i][0 : len(fileContents[i]) / 2]
        secondHalf := fileContents[i][len(fileContents[i]) / 2 : len(fileContents[i])]


        for j := 0; j < len(firstHalf); j++ {
            charFromFirstHalf := firstHalf[j]

            for k := 0; k < len(secondHalf); k++ {
                if charFromFirstHalf == secondHalf[k] {
                        initialItems = append(initialItems, charFromFirstHalf)
                }
            }
        }
        firstHalf = ""
        secondHalf = ""
    }

    finalItems := make([]byte, 0)
    checkForDuplicates := make(map[byte]int)
    for _, val := range initialItems {
        checkForDuplicates[val] = 1
    }

    for char := range checkForDuplicates {
        finalItems = append(finalItems, char)
    } 

    /* letter -> ASCII value -> points
        a -> 97 -> 1
        z -> 122 -> 26
        A -> 65 -> 27
        Z -> 90 -> 52
    */
    
    points := 0
    for _, val := range finalItems {
        if val <= 122 && val >= 97 {
            points += 1 + int(val) - 97
        } else {
            points += 27 + int(val) - 65
        }
    }
    
    return points
}

func main() {
    fmt.Println(partOne(readFile("sample.txt")))
}
