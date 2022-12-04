package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

func partOne(fileContents []string) int {
    horizontalPosition := 0
    depth := 0

    for _, line := range fileContents {
        contents := strings.Split(line, " ")
        
        value, err := strconv.Atoi(contents[1]) 
        if err != nil {
            fmt.Println(err.Error())
        }

        switch contents[0] {
        case "forward": horizontalPosition += value
        case "up": depth -= value
        case "down": depth += value
        } 
    }
    return horizontalPosition * depth
}

func partTwo(fileContents []string) int {
    horizontalPosition := 0
    depth := 0
    aim := 0

    for _, line := range fileContents {
        contents := strings.Split(line, " ")
        
        value, err := strconv.Atoi(contents[1]) 
        if err != nil {
            fmt.Println(err.Error())
        }

        switch contents[0] {
        case "forward": 
            horizontalPosition += value
            depth += aim * value
        case "up":
            aim -= value
        case "down": 
            aim += value
        } 
    }
    return horizontalPosition * depth
}

func main() {
    fmt.Println(partOne(readFile("input.txt")))
    fmt.Println(partTwo(readFile("input.txt")))
}


