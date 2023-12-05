package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func checkIfSymbol(array *[][]string, i, j int) bool {
	if (*array)[i][j] != "." {
		if _, err := strconv.Atoi((*array)[i][j]); err != nil {
			return true
		}
	}
	return false
}

func checkIfSymbolstar(array *[][]string, i, j int) bool {
	if (*array)[i][j] == "*" {
		return true
	}
	return false
}

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
	var arr [][]string
	// Iterate over each line in the file
	rows, cols := 0, 0
	for scanner.Scan() {
		line1 := strings.Split(scanner.Text(), "")
		arr = append(arr, line1)
		cols = len(line1)
		rows++
	}
	fmt.Println(rows, cols)
	for i := 0; i < rows; i++ {
		j, jEnd, jStart := 0, 0, 0
		for j < cols {
			if _, err := strconv.Atoi(arr[i][j]); err == nil {
				jStart = j
				jEnd = j
				for j < cols {
					if _, err := strconv.Atoi(arr[i][j]); err == nil {
						jEnd = j
						j++
					} else {
						break
					}
				}
			} else {
				j++
				continue
			}
			isSymbol := false
			//fmt.Println(jStart, jEnd)
			for k := jStart - 1; k <= jEnd+1; k++ {
				// check upper row
				if i-1 >= 0 && k >= 0 && k < cols {
					if checkIfSymbol(&arr, i-1, k) {
						isSymbol = true
						break
					}
				}
				if i+1 < rows && k < cols && k >= 0 {
					if checkIfSymbol(&arr, i+1, k) {
						isSymbol = true
						break
					}
				}
				//check below row
			}

			if (jStart-1) >= 0 && checkIfSymbol(&arr, i, jStart-1) {
				isSymbol = true
			}

			if jEnd+1 < cols && checkIfSymbol(&arr, i, jEnd+1) {
				isSymbol = true
			}

			/*
				if i-1,j or i+1,j or i,j-1 or i,j+1 or i-1,j-1 or i+1,j+1 or i+1,j-1 or i-1,j+1
				is a symbol
					then consider as part
			*/
			if isSymbol {
				num, _ := strconv.Atoi(strings.Join(arr[i][jStart:jEnd+1], ""))
				sum += num
				//fmt.Println(arr[i][jStart:jEnd], num)
			}
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
	var arr [][]string
	// Iterate over each line in the file
	rows, cols := 0, 0
	for scanner.Scan() {
		line1 := strings.Split(scanner.Text(), "")
		arr = append(arr, line1)
		cols = len(line1)
		rows++
	}
	fmt.Println(rows, cols)
	mapstar := make(map[string][]int)
	for i := 0; i < rows; i++ {
		j, jEnd, jStart := 0, 0, 0
		for j < cols {
			if _, err := strconv.Atoi(arr[i][j]); err == nil {
				jStart = j
				jEnd = j
				for j < cols {
					if _, err := strconv.Atoi(arr[i][j]); err == nil {
						jEnd = j
						j++
					} else {
						break
					}
				}
			} else {
				j++
				continue
			}
			num, _ := strconv.Atoi(strings.Join(arr[i][jStart:jEnd+1], ""))
			//fmt.Println(jStart, jEnd)
			identified := false
			for k := jStart - 1; k <= jEnd+1; k++ {
				// check upper row
				if i-1 >= 0 && k >= 0 && k < cols {
					if checkIfSymbolstar(&arr, i-1, k) {
						mapstar[fmt.Sprint([]int{i - 1, k})] = append(mapstar[fmt.Sprint([]int{i - 1, k})], num)
						identified = true
						break
					}
				}
				if i+1 < rows && k < cols && k >= 0 {
					if checkIfSymbolstar(&arr, i+1, k) {
						mapstar[fmt.Sprint([]int{i + 1, k})] = append(mapstar[fmt.Sprint([]int{i + 1, k})], num)
						identified = true
						break
					}
				}
				//check below row
			}

			if !identified && (jStart-1) >= 0 && checkIfSymbolstar(&arr, i, jStart-1) {
				mapstar[fmt.Sprint([]int{i, jStart - 1})] = append(mapstar[fmt.Sprint([]int{i, jStart - 1})], num)
			}

			if !identified && jEnd+1 < cols && checkIfSymbolstar(&arr, i, jEnd+1) {
				mapstar[fmt.Sprint([]int{i, jEnd + 1})] = append(mapstar[fmt.Sprint([]int{i, jEnd + 1})], num)
			}
		}
	}
	fmt.Println(mapstar)
	for _, v := range mapstar {
		if len(v) == 2 {
			sum += v[0] * v[1]
		}
	}
	return sum, nil
}

func main() {
	filename := "input" // Replace with your file name

	// Call the function to read from the file
	s, err := readFile(filename)
	fmt.Println(s)
	fmt.Println("part2")
	s, err = part2("input2")
	fmt.Println(s)
	if err != nil {
		log.Fatal(err)
	}
}
