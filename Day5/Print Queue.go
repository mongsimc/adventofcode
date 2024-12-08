package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Pair struct {
	X, Y int
}

func main() {

	fmt.Println("Reading file")
	filename := "input.txt"
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	var update []string
	var pageList []Pair
	intSet := make(map[int]struct{})
	for scanner.Scan() {
		line := scanner.Text()
		re := regexp.MustCompile(`(\d+)\|(\d+)`)
		reComplete := regexp.MustCompile(`(\d+,)+\d+`)

		matches := re.FindStringSubmatch(line)
		if matches != nil {
			x, _ := strconv.Atoi(matches[1])
			y, _ := strconv.Atoi(matches[2])
			c := Pair{X: x, Y: y}

			if _, exists := intSet[x]; !exists {
				intSet[x] = struct{}{}
			}
			if _, exists := intSet[y]; !exists {
				intSet[y] = struct{}{}
			}

			pageList = append(pageList, c)
		}

		update = append(update, reComplete.FindAllString(line, -1)...)

	}

	sortedList := SortPage(intSet, pageList)

	fmt.Printf("Answer1: %v\n", pageList)
	fmt.Printf("Answer2: %v\n", update)
	fmt.Printf("Answer1: %v\n", intSet)
}

func SortPage(intMap map[int]struct{}, pagePair []int) []int {
	var sortedInt []int
	for key := range intMap {
		if len(sortedInt) < 1 {
			sortedInt = append(sortedInt, key)
		} else {
			for 
		}
	}
}

// func SortPage(list []string) []int {
// 	var sortedInt []int
// 	for i := range list {
// 		reInt := regexp.MustCompile(`\d+`)
// 		matches := reInt.FindAllString(i, -1)
// 		x, _ := strconv.Atoi(matches[0])
// 		y, _ := strconv.Atoi(matches[1])

// 		if len(sortedInt) < 1 {
// 			sortedInt = append(sortedInt, x, y)
// 		} else {
// 			for j := range sortedInt {
// 				if Compare(x, j) {
// 					sortedInt = append(sortedInt, x)
// 				}
// 			}
// 		}

// 	}

// }

// func Compare(x, y int) bool{

// 	if ()

// }

// func CheckUpdate(pair Pair, input []string) []int {

// }
