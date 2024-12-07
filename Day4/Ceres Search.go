package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Reading file")
	filename := "input.txt"
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Printf("Answer: %s\n", line)
	}

}
