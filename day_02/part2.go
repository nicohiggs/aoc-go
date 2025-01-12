// similar to part 1 but we add a function that if the original is_safe() check doesnt work
// then iteratively create a copy of the array, delete one element, and check again. Essentially
// brute force
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

func is_safe2(nums []int) bool {
  if is_safe(nums) {
    return true
  }
  for i := 0; i < len(nums); i++ {
    nums2 := make([]int, 0)
    nums2 = append(nums2, nums[:i]...)
    nums2 = append(nums2, nums[i+1:]...)
    if is_safe(nums2) {
      return true
    }
  }
  return false
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
    if is_safe2(nums) {
      result += 1
    }
  }

  fmt.Printf("%v\n", result)
}
