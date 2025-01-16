package main

import (
  "fmt"
  "bufio"
  "os"
  "strings"
  "strconv"
)

func main() {
  file, _ := os.Open("input.txt")
  scanner := bufio.NewScanner(file)
  rules := make(map[int][]int)
  page_orders := [][]int{}
  i := 0
  rules_parsed := false
  for scanner.Scan() {
    line := scanner.Text()
    if !rules_parsed {
      if len(line) > 1 {
        curr_rule := strings.Split(line, "|")
        num1, _ := strconv.Atoi(curr_rule[0])
        num2, _ := strconv.Atoi(curr_rule[1])
        rules[num1] = append(rules[num1], num2)
      } else {
        rules_parsed = true
      }
    } else {
      page_orders = append(page_orders, []int{})
      curr_order := strings.Split(line, ",")
      for _, num := range curr_order {
        num_int, _ := strconv.Atoi(num)
        page_orders[i] = append(page_orders[i], num_int)
      }
      i++
    }
  }

  result := 0
  for _, po := range page_orders {
    valid_po := true
    for j:= len(po) - 1; j > 0; j-- {
      nums_less := rules[po[j]]
      for k := j - 1; k >= 0; k-- {
        for _, num := range nums_less {
          if num == po[k] {
            valid_po = false
          }
        }
      }
    }
    if valid_po {
      result += po[len(po)/2]
    }
  }
  fmt.Printf("%v\n", result)
}
