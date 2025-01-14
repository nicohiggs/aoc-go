package main

import (
  "fmt"
  "os"
  "bufio"
  "unicode"
  "strconv"
)

func main() {
  file, _ := os.Open("input.txt")
  defer file.Close()
  scanner := bufio.NewScanner(file)
  do_state := true
  result := 0
  for scanner.Scan() {
    line := scanner.Text()
    mul_found := false
    num1 := 0
    num2 := 0
    for i := 0; i < len(line);{
      // do parse
      if i < len(line) - 4 {
        if line[i:i+4] == "do()" {
          do_state = true
          i = i + 4
        }
      }
      if i < len(line) - 7 {
        if line[i:i+7] == "don't()" {
          do_state = false
          i = i + 7
        }
      }
      // mul parse
      mul_found = false
      if i < len(line) - 8 {
        if line[i:i+4] == "mul(" {
          i = i + 4
          j := i
          if unicode.IsDigit(rune(line[j])) {
            j++
            for unicode.IsDigit(rune(line[j])) {
              j++
            }
            if line[j] == ',' {
              k := j + 1
              if unicode.IsDigit(rune(line[k])) {
                k++
                for unicode.IsDigit(rune(line[k])) {
                  k++
                }
                if line[k] == ')' {
                  mul_found = true
                  num1, _ = strconv.Atoi(line[i:j])
                  num2, _ = strconv.Atoi(line[j+1:k])
                }
              }
            }
          }
        }
      }
      if mul_found && do_state {
        result += num1 * num2
      } else {
        i++
      }
    }
  }
  fmt.Printf("%v\n", result)
}
