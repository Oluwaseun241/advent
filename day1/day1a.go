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
  maxCalorie := 0

  for sc.Scan() {
    snack, err := strconv.Atoi(sc.Text())
    calorie += snack
     
    if err != nil{
      if calorie > maxCalorie{
        maxCalorie = calorie
      }
      calorie = 0
    }
  }
  fmt.Println(maxCalorie)
}
