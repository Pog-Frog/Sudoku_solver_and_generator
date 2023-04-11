# Sudoku Solver

This is a simple Sudoku Solver implemented in Go. It uses a backtracking algorithm to solve the Sudoku puzzles.

## Usage

To use the program, simply run the following command:

```bash
go run main.go
```

This will generate a random Sudoku puzzle, solve it, and print it out to the console.

## Customizing the Puzzle

You can customize the number of cells to remove from the generated puzzle by changing the `numToRemove` variable in the `makePuzzle` function.
You can also provide your own puzzle by modifying the `grid` variable in the `main` function.

## Input

If you want to play the game yourself, the program allows you to update the grid by accepting user input. Simply enter the row, column, and value (separated by spaces) when prompted.

## Output

The program will print out the Sudoku puzzle and the solution to the console.
