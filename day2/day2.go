package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func readFile(filename string, r, g, b int) (int, error) {
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
		gameID, err := strconv.Atoi(strings.Split(strings.Split(line, ":")[0], " ")[1])
		if err != nil {
			fmt.Println(err)
		}
		isValid := true
		tries := strings.Split(strings.Split(line, ":")[1], ";")
		for _, t := range tries {
			cubes := strings.Split(t, ",")
			for _, c := range cubes {
				nStr, color := strings.Fields(c)[0], strings.Fields(c)[1]
				n, _ := strconv.Atoi(nStr)
				switch color {
				case "red":
					isValid = n <= r
					break
				case "green":
					isValid = n <= g
					break
				case "blue":
					isValid = n <= b
					break
				}
				if !isValid {
					break
				}
			}
			if !isValid {
				break
			}
		}
		if isValid {
			sum += gameID
		}
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return sum, nil
}

func part2(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		tries := strings.Split(strings.Split(line, ":")[1], ";")
		r, g, b := -1, -1, -1
		for _, t := range tries {
			cubes := strings.Split(t, ",")
			for _, c := range cubes {
				nStr, color := strings.Fields(c)[0], strings.Fields(c)[1]
				n, _ := strconv.Atoi(nStr)
				switch color {
				case "red":
					r = int(math.Max(float64(r), float64(n)))
					break
				case "green":
					g = int(math.Max(float64(g), float64(n)))
					break
				case "blue":
					b = int(math.Max(float64(b), float64(n)))
					break
				}
			}
		}
		sum += r * g * b
	}
	return sum, nil
}

func main() {
	filename := "input" // Replace with your file name

	// Call the function to read from the file
	s, err := readFile(filename, 12, 13, 14)
	fmt.Println(s)
	fmt.Println("part2")
	s, err = part2("input2")
	fmt.Println(s)
	if err != nil {
		log.Fatal(err)
	}
}
