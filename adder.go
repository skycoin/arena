package main

import (
  "bufio"
  "fmt"
  "os"
  "strings"
  "strconv"
)

func main() {
  var correct = true
  reader := bufio.NewReader(os.Stdin)
  fmt.Println("Input digit 1")
  fmt.Print("--->")
  text1, _ := reader.ReadString('\n')
  text1 = strings.Replace(text1, "\n", "", -1)
  digit1, err := strconv.ParseFloat(text1,64)
	if err != nil {
		  correct = false
	}

  fmt.Println("Input digit 2")
  fmt.Print("--->")
  
  text2, _ := reader.ReadString('\n')
  text2 = strings.Replace(text2, "\n", "", -1)

  digit2, err := strconv.ParseFloat(text2,64)
	if err != nil {
		  correct = false
	}
  if correct{
      fmt.Println(digit1+digit2)
  } else{
      fmt.Println("Wrong input!")
  }

}
