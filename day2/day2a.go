package main

import (
	"bufio"
	"fmt"
	"os"
)

// A : Rock
// B : Paper
// C : Scissors

func strategy(input map[string]string) int {
  score := 0
  for _, response := range input {
    if response == "Y" {
      score += 8
    } else if response == "X" {
      score += 1
    } else if response == "Z" {
      score += 6
    }
  }
  return score
}

func main() {
  input,_ := os.Open("input.txt")
  defer input.Close()
  sc := bufio.NewScanner(input)

  for sc.Scan() {
    score := sc.Text()
  }

  totalScore := strategy(input)
  fmt.Println(totalScore)
}
