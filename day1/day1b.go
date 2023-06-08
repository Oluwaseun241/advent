package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)


func main() {
  input,_ := os.Open("input.txt")
  defer input.Close()
  sc := bufio.NewScanner(input)

  calorie := 0
  maxCalorie1 := 0
  maxCalorie2 := 0
  maxCalorie3 := 0

  for sc.Scan() {
    snack, err := strconv.Atoi(sc.Text())
    calorie += snack 

    if err != nil {
      if calorie > maxCalorie3 {
        maxCalorie1 = maxCalorie2
        maxCalorie2 = maxCalorie3
        maxCalorie3 = calorie
      } else if calorie > maxCalorie2 {
        maxCalorie1 = maxCalorie2
        maxCalorie2 = calorie
      } else if calorie > maxCalorie3 {
        maxCalorie1 = calorie
      } 
      calorie = 0 
    }
  }
  totalCalorie := maxCalorie1+maxCalorie2+maxCalorie3
  fmt.Println(totalCalorie)
}
