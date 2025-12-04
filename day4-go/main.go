package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func getMatrix(fichName string) [][]int {
	f, err := os.Open(fichName)
	if err != nil {
		log.Fatal(err)
	}
	defer func(fs os.File) {
		if err := fs.Close(); err != nil {
			log.Print(err)
		}
	}(*f)

	scanner := bufio.NewScanner(f)
	matrix := make([][]int, 0)

	for scanner.Scan() {
		line := scanner.Text()
		matrLine := make([]int, 0)
		for _, c := range line {
			switch c {
			case '.':
				matrLine = append(matrLine, 0)
			case '@':
				matrLine = append(matrLine, 1)
			}
		}
		matrix = append(matrix, matrLine)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return matrix
}

func main() {
	fmt.Print("Do you want to execute ex 1 or 2? [1,2]: ")
	var opt int
	_, err := fmt.Scanln(&opt)
	if err != nil {
		log.Fatal("Error reading input")
	}

	switch opt {
	case 1:
		ex1()
	case 2:
		ex2()
	default:
		fmt.Println("Please pick 1 or 2")
	}
}
