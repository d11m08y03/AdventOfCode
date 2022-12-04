package main

import (
	"bufio"
	"fmt"
	"os"
    "strconv"
)

func findMax(numbers []int) int {
    max := numbers[0]

    for _, value := range numbers {
        if value > max {
            max = value
        }
    }

    return max
}

func findSumTopThree(numbers []int) int {
    first := findMax(numbers)

    second := numbers[0]
    for _, value := range numbers {
        if value > second && value != first {
            second = value
        }         
    }

    third := numbers[0]
    for _, value := range numbers {
        if value > third && value != first && value != second {
            third = value
        }
    }

    return first + second + third
}

func main()  {
    file, err := os.Open("input.txt")
    if err != nil {
        fmt.Println(err.Error())
    }
    defer file.Close()

    var numbers []int
    var currLine string
    currSum := 0
    var currNum int
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        currLine = scanner.Text()
        if currLine != "" {
            currNum, err = strconv.Atoi(currLine)
            if err != nil {
                fmt.Println(err.Error())
            } else {
                currSum += currNum
            }
        } else {
            numbers = append(numbers, currSum)
            currSum = 0
        }
    }

    fmt.Printf("Highest calories carried: %v\n", findMax(numbers))
    fmt.Printf("Total carried by top 3: %v\n", findSumTopThree(numbers))

}


