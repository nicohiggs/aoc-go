// parse into an array of ints and then iteratively check that the array
// follows the 'monotonic' rules with the is_safe() function
package main

import (
  "os"
  "fmt"
  "bufio"
  "strings"
  "strconv"
)

func is_safe(nums []int) bool {
  increasing := true
  if nums[1] < nums[0] {
    increasing = false
  }
  for i := 1; i < len(nums); i++ {
    curr := nums[i]
    prev := nums[i-1]
    if increasing {
      if curr <= prev {
        return false
      }
      if (curr - prev) > 3 {
        return false
      }
    } else {
      if prev <= curr {
        return false
      }
      if (prev - curr) > 3 {
        return false
      }
    }
  }
  return true
}

func main() {
  file, err := os.Open("input.txt")
  if err != nil {
    fmt.Print("Failed opening file.")
  }
  defer file.Close()

  result := 0
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    line := scanner.Text()
    nums_str := strings.Split(line, " ")
    nums := []int{}
    for _, v := range nums_str {
      num, err := strconv.Atoi(v)
      if err != nil {
        fmt.Print("failed to conv string to int")
      }
      nums = append(nums, num)
    }
    if is_safe(nums) {
      result += 1
    }
  }

  fmt.Printf("%v\n", result)
}
