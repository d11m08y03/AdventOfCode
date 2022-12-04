package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func readFile(filename string) []int {
    var fileContents []int 

    file, err := os.Open(filename)
    if err != nil {
        println(err.Error())
    }
    defer file.Close()
    
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        value, err := strconv.Atoi(scanner.Text())
        if err != nil {
            fmt.Println(err.Error())
        }
        fileContents = append(fileContents, value)
    }
    return fileContents
}

func calculateIncreasesPartOne(fileContents []int) int {
    count := 0
    
    for i := 1; i < len(fileContents); i++ {
        if fileContents[i] - fileContents[i - 1] > 0 {
            count += 1
        }
    }

    return count 
}

func calculateIncreasesPartTwo(fileContents []int) int {
    count := 0

    for i := 0; i < len(fileContents) - 3; i++ {
        // if (fileContents[i] + fileContents[i + 1] + fileContents[i + 2]) < (fileContents[i + 1] + fileContents[i + 2] + fileContents[i + 3]) {
        //     count += 1
        // }

        // the math do be mathing 
        if fileContents[i] < fileContents[i + 3] {
            count += 1
        }
    }

    return count
}

func main() {
    fmt.Println(calculateIncreasesPartOne(readFile("input.txt")))
    fmt.Println(calculateIncreasesPartTwo(readFile("input.txt")))

}
