<!-- # Sudoku Solver

## Instructions

Create a program that resolves a sudoku. A valid sudoku has only one possible solution. Make sure you submit all the necessary files to run the program.

## Usage

### Example 1:

Example of output for a valid sudoku:

```bash
$ go run . ".96.4...1" "1...6...4" "5.481.39." "..795..43" ".3..8...." "4.5.23.18" ".1.63..59" ".59.7.83." "..359...7" | cat -e
3 9 6 2 4 5 7 8 1$
1 7 8 3 6 9 5 2 4$
5 2 4 8 1 7 3 9 6$
2 8 7 9 5 1 6 4 3$
9 3 1 4 8 6 2 7 5$
4 6 5 7 2 3 9 1 8$
7 1 2 6 3 8 4 5 9$
6 5 9 1 7 4 8 3 2$
8 4 3 5 9 2 1 6 7$
$
```

### Example 2:

Examples of expected outputs for invalid inputs or sudokus:

```bash
$ go run . 1 2 3 4 | cat -e
Error$
$ go run . | cat -e
Error$
$ go run . ".96.4...1" "1...6.1.4" "5.481.39." "..795..43" ".3..8...." "4.5.23.18" ".1.63..59" ".59.7.83." "..359...7" | cat -e
Error$
$
``` -->


# Sudoku Solver

This project is a Sudoku game solver created with the aim of providing a solution for valid Sudoku puzzles. Sudoku is a classic number puzzle where you need to fill a 9x9 grid with digits so that each column, each row, and each of the nine 3x3 subgrids that compose the grid (also known as "boxes") contain all of the digits from 1 to 9.

## Getting Started

To use this Sudoku Solver, follow the instructions below to compile and run the program. Make sure you have Go installed on your system.

### Prerequisites

- Go (Golang) is required to compile and run this program.

### Installation

1. Clone this repository to your local machine or download the project files.

2. Open your terminal and navigate to the main project directory.

3. Compile the program using the `go run` command:

```bash
go run . "<Sudoku_Row_1>" "<Sudoku_Row_2>" "<Sudoku_Row_3>" ... "<Sudoku_Row_9>" | cat -e
```

Replace `<Sudoku_Row_1>` to `<Sudoku_Row_9>` with your Sudoku puzzle. Use dots (`.`) to represent empty cells, and separate the rows with spaces.

For example, to solve a Sudoku puzzle, run:

```bash
go run . ".96.4...1" "1...6...4" "5.481.39." "..795..43" ".3..8...." "4.5.23.18" ".1.63..59" ".59.7.83." "..359...7" | cat -e
```

The program will output the solution if the Sudoku is valid. If there's no solution or the input is invalid, it will display "Error."

### Examples

Here are some usage examples:

#### Valid Sudoku Example:

Input:
```
$ go run . ".96.4...1" "1...6...4" "5.481.39." "..795..43" ".3..8...." "4.5.23.18" ".1.63..59" ".59.7.83." "..359...7" | cat -e
```

Output:
```
3 9 6 2 4 5 7 8 1$
1 7 8 3 6 9 5 2 4$
5 2 4 8 1 7 3 9 6$
2 8 7 9 5 1 6 4 3$
9 3 1 4 8 6 2 7 5$
4 6 5 7 2 3 9 1 8$
7 1 2 6 3 8 4 5 9$
6 5 9 1 7 4 8 3 2$
8 4 3 5 9 2 1 6 7$
```

#### Invalid Input Examples:

1. Invalid number of arguments:
   Input:
   ```
   $ go run . 1 2 3 4 | cat -e
   ```

   Output:
   ```
   Error$
   ```

2. No input provided:
   Input:
   ```
   $ go run . | cat -e
   ```

   Output:
   ```
   Error$
   ```

3. Invalid Sudoku (conflicting numbers in the same row):
   Input:
   ```
   $ go run . ".96.4...1" "1...6.1.4" "5.481.39." "..795..43" ".3..8...." "4.5.23.18" ".1.63..59" ".59.7.83." "..359...7" | cat -e
   ```

   Output:
   ```
   Error$
   ```

## Acknowledgments

- Thanks to the developers of Go for providing a versatile programming language.
- Sudoku is a classic puzzle, and this project is an attempt to solve it programmatically.

---

*Note: This README file provides an overview of the Sudoku Solver project. For more detailed information, refer to the code and documentation in the project's source files.*