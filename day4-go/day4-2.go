package main

import (
	"fmt"
)

func ex2() {
	paperN := 0
	paperAnt := -1
	matrix := getMatrix("day4.txt")

	for _, line := range matrix {
		for _, num := range line {
			fmt.Print(num)
		}
		fmt.Println()
	}

	for paperAnt != paperN {
		paperAnt = paperN

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
						matrix[row][col] = 0
					}
				}
			}
		}
	}

	fmt.Println()
	fmt.Println("The paper rolls that can be accessed are:", paperN)
}
