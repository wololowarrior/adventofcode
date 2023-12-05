package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func readFile(filename string) (int, error) {
	// Open the file for reading
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	sum := 0
	// Iterate over each line in the file
	for scanner.Scan() {
		line := scanner.Text()
		f, l := -1, -1
		for i := 0; i < len(line); i++ {
			t, err := strconv.Atoi(string(line[i]))
			if err == nil && f == -1 {
				f = t
			}

			t1, err1 := strconv.Atoi(string(line[len(line)-i-1]))
			if err1 == nil && l == -1 {
				l = t1
			}
			if f > -1 && l > -1 {
				break
			}
		}
		sum += f*10 + l
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return sum, nil
}

func main() {
	filename := "input" // Replace with your file name

	// Call the function to read from the file
	s, err := readFile(filename)
	fmt.Println(s)
	if err != nil {
		log.Fatal(err)
	}
}
