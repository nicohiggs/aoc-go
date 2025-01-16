/*
The approach is to parse the input into a 2-d grid, then iterate through and at each
'X'check the 8 directions for an XMAS match
*/
package main

import (
  "fmt"
  "bufio"
  "os"
)

func main() {
  file, _ := os.Open("input.txt")
  scanner := bufio.NewScanner(file)
  grid := [][]rune{}
  i := 0
  for scanner.Scan() {
    line := scanner.Text()
    grid = append(grid, []rune{})
    for _, c := range line {
      grid[i] = append(grid[i], c)
    }
    i++
  }
  result := 0
  for i, row := range grid {
    for j, c := range row {
      if c == 'X' {
        // up
        if i > 2 && grid[i-1][j] == 'M' && grid[i-2][j] == 'A' && grid[i-3][j] == 'S' {
          result++
        }
        // right
        if j < len(row) - 3 && grid[i][j+1] == 'M' && grid[i][j+2] == 'A' && grid[i][j+3] == 'S' {
          result++
        }
        // down
        if i < len(grid) - 3 && grid[i+1][j] == 'M' && grid[i+2][j] == 'A' && grid[i+3][j] == 'S' {
          result++
        }
        // left
        if j > 2 && grid[i][j-1] == 'M' && grid[i][j-2] == 'A' && grid[i][j-3] == 'S' {
          result++
        }
        // up-left
        if i > 2 && j > 2 && grid[i-1][j-1] == 'M' && grid[i-2][j-2] == 'A' && grid[i-3][j-3] == 'S' {
          result++
        }
        // up-right
        if i > 2 && j < len(row) - 3 && grid[i-1][j+1] == 'M' && grid[i-2][j+2] == 'A' && grid[i-3][j+3] == 'S' {
          result++
        }
        // down-left
        if i < len(grid) - 3 && j > 2 && grid[i+1][j-1] == 'M' && grid[i+2][j-2] == 'A' && grid[i+3][j-3] == 'S' {
          result++
        }
        // down-right
        if i < len(grid) - 3 && j < len(row) - 3 && grid[i+1][j+1] == 'M' && grid[i+2][j+2] == 'A' && grid[i+3][j+3] == 'S' {
          result++
        }
      }
    }
  }
  fmt.Printf("%v\n", result)
}
