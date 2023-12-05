package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
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
	for scanner.Scan() {
		line := scanner.Text()
		list := strings.Split(strings.Split(line, ":")[1], "|")
		cardVal := 1
		isvalulable := true
		mapofwin := make(map[string]bool)

		for _, v := range strings.Fields(list[0]) {
			//i, _ := strconv.Atoi(v)
			mapofwin[v] = true
		}

		for _, v := range strings.Fields(list[1]) {
			if _, ok := mapofwin[v]; ok {
				cardVal = cardVal << 1
				isvalulable = true
			}
		}
		if isvalulable {
			sum += cardVal / 2
		}
	}
	return sum, nil
}

func part2(filename string) (int, error) {
	// Open the file for reading
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	sum := 0
	cardN := 1
	mapres := make(map[int]int)
	for scanner.Scan() {
		line := scanner.Text()
		list := strings.Split(strings.Split(line, ":")[1], "|")
		cardmatched := 0
		mapofwin := make(map[string]bool)
		if _, ok := mapres[cardN]; !ok {
			mapres[cardN] = 1
		}
		for _, v := range strings.Fields(list[0]) {
			//i, _ := strconv.Atoi(v)
			mapofwin[v] = true
		}

		for _, v := range strings.Fields(list[1]) {
			if _, ok := mapofwin[v]; ok {
				cardmatched++
			}
		}
		fmt.Println(cardN, cardmatched)
		for i := cardN + 1; i <= cardN+cardmatched; i++ {
			if _, ok := mapres[i]; !ok {
				mapres[i] = 1
			}
			mapres[i] = mapres[i] + mapres[cardN]
		}

		cardN++
	}
	fmt.Println(mapres)
	for _, v := range mapres {
		sum += v
	}
	return sum, nil
}

func main() {
	filename := "input" // Replace with your file name

	// Call the function to read from the file
	s, err := readFile(filename)
	fmt.Println(s)
	fmt.Println("part2")
	s, err = part2("input")
	fmt.Println(s)
	if err != nil {
		log.Fatal(err)
	}
}
