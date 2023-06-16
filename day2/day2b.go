package main

import(
  "fmt"
)

func strategy(input map[string]string) int {
  score := 0
  for _, response := range input {
    if response == "Y" {
      score += 4
    } else if response == "X" {
      score += 1
    } else if response == "Z" {
      score += 7
    }
  }
  return score
}

func main() {
input := map[string]string {
     "A": "Y",
     "B": "X",
     "C": "Z",
  }

  totalScore := strategy(input)
  fmt.Println(totalScore)
}
