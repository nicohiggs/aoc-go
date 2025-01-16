
/*
The approach is to parse the input into a 2-d grid, then iterate through and at each
'A'check for an X-MAS match
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
      if c == 'A' {
        if i > 0 && i < len(grid) - 1 && j > 0 && j < len(row) - 1 {
          if ((grid[i-1][j-1] == 'M' && grid[i+1][j+1] == 'S') || (grid[i-1][j-1] == 'S' && grid[i+1][j+1] == 'M')) &&
             ((grid[i-1][j+1] == 'M' && grid[i+1][j-1] == 'S') || (grid[i-1][j+1] == 'S' && grid[i+1][j-1] == 'M')) {
            result++
          }
        }
      }
    }
  }
  fmt.Printf("%v\n", result)
}
