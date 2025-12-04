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
	paperN := 0
	matrix := getMatrix("day4-example.txt")

	for _, line := range matrix {
		for _, num := range line {
			print(num)
		}
		println()
	}

	println()
	fmt.Println("The paper rolls that can be accessed are:", paperN)
}
