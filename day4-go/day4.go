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
	matrix := getMatrix("day4.txt")

	for _, line := range matrix {
		for _, num := range line {
			fmt.Print(num)
		}
		fmt.Println()
	}

	for row, rowElem := range matrix {
		for col := range rowElem {
			cont := 0

			if matrix[row][col] == 1 {
				switch row {
				case 0:
					if matrix[row+1][col] == 1 {
						cont++
					}
					switch col {
					case 0:
						if matrix[row+1][col+1] == 1 {
							cont++
						}
					case len(rowElem) - 1:
						if matrix[row+1][col-1] == 1 {
							cont++
						}
					default:
						if matrix[row+1][col+1] == 1 {
							cont++
						}
						if matrix[row+1][col-1] == 1 {
							cont++
						}
					}
				case len(matrix) - 1:
					if matrix[row-1][col] == 1 {
						cont++
					}
					switch col {
					case 0:
						if matrix[row-1][col+1] == 1 {
							cont++
						}
					case len(rowElem) - 1:
						if matrix[row-1][col-1] == 1 {
							cont++
						}
					default:
						if matrix[row-1][col+1] == 1 {
							cont++
						}
						if matrix[row-1][col-1] == 1 {
							cont++
						}
					}
				default:
					if matrix[row+1][col] == 1 {
						cont++
					}
					if matrix[row-1][col] == 1 {
						cont++
					}
					switch col {
					case 0:
						if matrix[row+1][col+1] == 1 {
							cont++
						}
						if matrix[row-1][col+1] == 1 {
							cont++
						}
					case len(rowElem) - 1:
						if matrix[row+1][col-1] == 1 {
							cont++
						}
						if matrix[row-1][col-1] == 1 {
							cont++
						}
					default:
						if matrix[row+1][col+1] == 1 {
							cont++
						}
						if matrix[row+1][col-1] == 1 {
							cont++
						}
						if matrix[row-1][col+1] == 1 {
							cont++
						}
						if matrix[row-1][col-1] == 1 {
							cont++
						}
					}
				}

				switch col {
				case 0:
					if matrix[row][col+1] == 1 {
						cont++
					}
				case len(rowElem) - 1:
					if matrix[row][col-1] == 1 {
						cont++
					}
				default:
					if matrix[row][col+1] == 1 {
						cont++
					}
					if matrix[row][col-1] == 1 {
						cont++
					}
				}

				if cont < 4 {
					paperN++
				}
			}
		}
	}

	fmt.Println()
	fmt.Println("The paper rolls that can be accessed are:", paperN)
}
