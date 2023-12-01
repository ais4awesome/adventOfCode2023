// PART ONE

// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strconv"
// 	"strings"
// )

// func main() {
// 	var inputList []string
// 	var calibrationList []int

// 	file, err := os.Open("input.txt")
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	defer file.Close()

// 	scanner := bufio.NewScanner(file)
// 	for scanner.Scan() {
// 		inputList = append(inputList, scanner.Text())
// 	}

// 	if err := scanner.Err(); err != nil {
// 		fmt.Println(err)
// 	}

// 	// this wont easily work due to indexing as a rune, not a byte
// 	// Besides the axiomatic detail that Go source code is UTF-8, there’s really only one way
// 	// that Go treats UTF-8 specially, and that is when using a for range loop on a string.
// 	// We’ve seen what happens with a regular for loop. A for range loop, by contrast,
// 	// decodes one UTF-8-encoded rune on each iteration.
// 	// for _, line := range inputList {
// 	// 	// input is good
// 	// 	//println(line)
// 	// 	var intList []int
// 	// 	for _, char := range line {
// 	// 		// this is a rune
// 	// 		// println(char)
// 	// 		// // notice how we must conv runes to stings
// 	// 		// println(string(char))
// 	// 		// os.Exit(0)
// 	// 		if string(char) > "0" || string(char) < "9" {
// 	// 			intList = append(intList, int(char-'0'))
// 	// 		}
// 	// 	}
// 	// 	calibrationList = append(calibrationList, intList[0])
// 	// 	calibrationList = append(calibrationList, intList[len(intList)-1])
// 	// }

// 	for _, line := range inputList {
// 		var valuesList []string
// 		for _, char := range strings.Split(line, "") {
// 			_, err := strconv.Atoi(char)
// 			// only add to list when int
// 			if err == nil {
// 				valuesList = append(valuesList, char)
// 			}

// 		}
// 		// make assumption that list will always have at least one number per line
// 		if len(valuesList) < 2 {
// 			// add twice if only one number in list
// 			strVal := valuesList[0]
// 			strVal = strVal + strVal
// 			intVal, err := strconv.Atoi(strVal)
// 			if err != nil {
// 				fmt.Println(err.Error())
// 			}
// 			calibrationList = append(calibrationList, intVal)
// 		} else {
// 			// first val
// 			strVal := valuesList[0]
// 			// last val
// 			strVal = strVal + valuesList[(len(valuesList)-1)]
// 			intVal, err := strconv.Atoi(strVal)
// 			if err != nil {
// 				fmt.Println(err.Error())
// 			}
// 			calibrationList = append(calibrationList, intVal)
// 		}
// 	}

// 	sum := 0
// 	for _, line := range calibrationList {
// 		fmt.Println(line)
// 		sum = sum + line
// 	}
// 	fmt.Println("Total: " + string(sum))
// }

// PART TWO
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var inputList []string
	var calibrationList []int

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		inputList = append(inputList, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	for _, line := range inputList {
		var valuesList []string
		//var builder strings.Builder
		// instructions dont cover edge case for order of operations... may need to change -1
		if strings.Contains(line, "one") {
			line = strings.Replace(line, "one", "o1e", -1)
		}
		if strings.Contains(line, "two") {
			line = strings.Replace(line, "two", "t2o", -1)
		}
		if strings.Contains(line, "three") {
			line = strings.Replace(line, "three", "th3ee", -1)
		}
		if strings.Contains(line, "four") {
			line = strings.Replace(line, "four", "fo4r", -1)
		}
		if strings.Contains(line, "five") {
			line = strings.Replace(line, "five", "fi5e", -1)
		}
		if strings.Contains(line, "six") {
			line = strings.Replace(line, "six", "s6x", -1)
		}
		if strings.Contains(line, "seven") {
			line = strings.Replace(line, "seven", "se7en", -1)
		}
		if strings.Contains(line, "eight") {
			line = strings.Replace(line, "eight", "ei8ht", -1)
		}
		if strings.Contains(line, "nine") {
			line = strings.Replace(line, "nine", "ni9e", -1)
		}
		println(line)
		for _, char := range strings.Split(line, "") {
			_, err := strconv.Atoi(char)
			// only add to list when int
			if err == nil {
				valuesList = append(valuesList, char)
			}
		}
		//make assumption that list will always have at least one number per line
		if len(valuesList) < 2 {
			// add twice if only one number in list
			strVal := valuesList[0]
			strVal = strVal + strVal
			intVal, err := strconv.Atoi(strVal)
			if err != nil {
				fmt.Println(err.Error())
			}
			calibrationList = append(calibrationList, intVal)
		} else {
			// first val
			strVal := valuesList[0]
			// last val
			strVal = strVal + valuesList[(len(valuesList)-1)]
			intVal, err := strconv.Atoi(strVal)
			if err != nil {
				fmt.Println(err.Error())
			}
			calibrationList = append(calibrationList, intVal)
		}
	}

	sum := 0
	for _, line := range calibrationList {
		//fmt.Println(line)
		sum = sum + line
	}
	fmt.Println("Total: ", sum)
}

// 94998 too large
//55358
