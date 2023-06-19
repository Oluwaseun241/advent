package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
  f,_ := os.Open("input.txt")
  defer f.Close()
  sc := bufio.NewScanner(f)

  scores := map[string]int {
    "A Y":4,
    "A X":3,
    "A Z":8,
    "B X":1,
    "B Y":5,
    "B Z":9,
    "C Z":7,
    "C Y":6,
    "C X":2,
  }

  var score int
  for sc.Scan() {
    score += scores[sc.Text()]
  }
  fmt.Println(score)
}

// func strategy(input map[string]string) int {
//   score := 0
//   for _, response := range input {
//     if response == "Y" {
//       score += 4
//     } else if response == "X" {
//       score += 1
//     } else if response == "Z" {
//       score += 7
//     }
//   }
//   return score
// }
//
// func main() {
// input := map[string]string {
//      "A": "Y",
//      "B": "X",
//      "C": "Z",
//   }
//
//   totalScore := strategy(input)
//   fmt.Println(totalScore)
// }
