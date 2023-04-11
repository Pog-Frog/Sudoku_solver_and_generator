package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	playGame()

}

func printPuzzle(grid [][]int) {
	fmt.Println("  1 2 3  4 5 6  7 8 9 ")
	for i := 0; i < 9; i++ {
		if i%3 == 0 {
			fmt.Println(" ---------------------")
		}
		fmt.Printf("%d", i+1)
		for j := 0; j < 9; j++ {
			if j%3 == 0 {
				fmt.Printf("|")
			}
			fmt.Printf("%d ", grid[i][j])
		}
		fmt.Printf("|\n")
	}
	fmt.Println(" ---------------------")
}

func checkRow(grid [][]int, row int, col int, num int) bool {
	for i := 0; i < 9; i++ {
		if grid[row][i] == num && grid[row][i] != 0 {
			return false
		}
	}
	return true
}

func checkCol(grid [][]int, row int, col int, num int) bool {
	for i := 0; i < 9; i++ {
		if grid[i][col] == num && grid[i][col] != 0 {
			return false
		}
	}
	return true
}

func checkBox(grid [][]int, row int, col int, num int) bool {
	rowStart := row - row%3
	colStart := col - col%3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if grid[rowStart+i][colStart+j] == num && grid[rowStart+i][colStart+j] != 0 {
				return false
			}
		}
	}
	return true
}

func checkSafe(grid [][]int, row int, col int, num int) bool {
	return checkRow(grid, row, col, num) && checkCol(grid, row, col, num) && checkBox(grid, row, col, num)
}

func solveSudoku(grid [][]int) bool {
	var row, col int
	found := false
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if grid[i][j] == 0 {
				row, col = i, j
				found = true
				break
			}
		}
		if found {
			break
		}
	}
	if !found {
		return true
	}
	for i := 1; i <= 9; i++ {
		if checkSafe(grid, row, col, i) {
			grid[row][col] = i
			if solveSudoku(grid) {
				return true
			}
			grid[row][col] = 0
		}
	}
	return false
}

func generatePuzzle() [][]int {
	allowedNums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	grid := make([][]int, 9)
	seed := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(seed)
	rng.Shuffle(len(allowedNums), func(i, j int) { allowedNums[i], allowedNums[j] = allowedNums[j], allowedNums[i] })
	grid[0] = allowedNums
	for i := 1; i < 9; i++ {
		if i%3 == 0 {
			grid[i] = shiftRow(grid[i-1], 1)
		} else {
			grid[i] = shiftRow(grid[i-1], 3)
		}
	}

	if solveSudoku(grid) {
		return grid
	} else {
		return generatePuzzle()
	}
}

func shiftRow(row []int, shift int) []int {
	for i := 0; i < shift; i++ {
		row = append(row[1:], row[0])
	}
	return row
}

func makePuzzle(grid [][]int, numToRemove int) [][]int {
	seed := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(seed)
	for i := 0; i < numToRemove; i++ {
		row := rng.Intn(9)
		col := rng.Intn(9)
		if grid[row][col] != 0 {
			grid[row][col] = 0
		} else {
			i--
		}
	}
	return grid
}

func getUserInput() (int, int, int, error) {
	var row, col, val int
	fmt.Println("Enter row, column and value (separated by space):")
	_, err := fmt.Scanf("%d %d %d", &row, &col, &val)
	if err != nil {
		return 0, 0, 0, err
	}
	if row < 1 || row > 9 || col < 1 || col > 9 || val < 1 || val > 9 {
		fmt.Println("Invalid input: row, column, and value must be between 1 and 9")
		return 0, 0, 0, fmt.Errorf("invalid input: row, column, and value must be between 1 and 9")
	}
	return row, col, val, nil
}

func updateGrid(grid [][]int, row int, col int, val int) bool {
	println("\nUpdating grid...")
	if checkSafe(grid, row, col, val) {
		grid[row][col] = val
		return true
	} else {
		fmt.Println("Invalid move! Please try again.")
		return false
	}
}

func checkSolved(grid [][]int) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if grid[i][j] == 0 {
				return false
			}
		}
	}
	return true
}

func playGame() {

	var grid [][]int = generatePuzzle()
	var level int
	fmt.Println("Choose difficulty level: 1 (easy), 2 (medium), 3 (hard)")
	_, err := fmt.Scanln(&level)
	if err != nil {
		return
	}

	var difficulty int
	switch level {
	case 1:
		difficulty = 30
	case 2:
		difficulty = 40
	case 3:
		difficulty = 50
	default:
		fmt.Println("Invalid choice, generating puzzle of medium difficulty")
		difficulty = 40
	}
	var generatedPuzzle [][]int = makePuzzle(grid, difficulty)
	var playerGrid [][]int = make([][]int, 9)
	for i := 0; i < 9; i++ {
		playerGrid[i] = make([]int, 9)
		copy(playerGrid[i], generatedPuzzle[i])
	}

	for {
		fmt.Println("\n\n")
		printPuzzle(playerGrid)
		row, col, val, _ := getUserInput()
		print("row: ", row, " col: ", col, " val: ", val, "")
		if row == 0 && col == 0 && val == 0 {
			continue
		}
		if updateGrid(playerGrid, row-1, col-1, val) {
			playerGrid[row-1][col-1] = val
			if checkSolved(playerGrid) {
				fmt.Println("Congratulations! You solved the puzzle!")
				return
			}
		}
	}
}
