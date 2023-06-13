package main

import (
	"bufio"
	"fmt"
	"os"
)

// A : Rock
// B : Paper
// C : Scissors



func main() {
  f,_ := os.Open("input.txt")
  defer f.Close()
  sc := bufio.NewScanner(f)
  
  scores := map[string]int {
    "A Y":8,
    "A X":4,
    "A Z":3,
    "B X":1,
    "B Y":5,
    "B Z":9,
    "C Z":6,
    "C Y":2,
    "C X":7,
  }

  var score int

  for sc.Scan() {
    score += scores[sc.Text()]
  }
  fmt.Println(score)
}

// SOlUTION TO THE EXAMPLE INPUT

// func strategy(input map[string]string) int {
//   score := 0
//   for _, response := range input {
//     if response == "Y" {
//       score += 8
//     } else if response == "X" {
//       score += 1
//     } else if response == "Z" {
//       score += 6
//     }
//   }
//   return score
// }

// func main() {
// input := map[string]string {
//      "A": "Y",
//      "B": "X",
//      "C": "Z",
//   }
//   fmt.Println(input)
//
//   totalScore := strategy(input)
//   fmt.Println(totalScore)
// }

