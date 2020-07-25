package main
import (
	"fmt"
	"os"
	"strconv"
	"errors"
)

func ErrorFormatter(ErrorMessage string) error {
	return errors.New(ErrorMessage)
}

func CheckForInValidInputSize(arr []string) bool {
	return len(arr) != 2
}

func addTwoNumbers(NumberArray []string) (float64,error) {
	var i int
	var sum float64

	for i = 0; i < len(NumberArray);i++ {
		var number, err = strconv.ParseFloat(NumberArray[i], 64)
		if err != nil{
			return sum,ErrorFormatter("Something went wrong")
		}
		sum += number
	}
	return sum, nil
}

func main() {
	var args = os.Args[1:]
	if CheckForInValidInputSize(args) {
		fmt.Println(ErrorFormatter("Invalid input size"))
		return
	}
	var err error
	var sum float64
	sum, err = addTwoNumbers(args)
	if err != nil {
		fmt.Println(ErrorFormatter("Invalid input"))
		return
	}
	fmt.Println(sum)
}


// Contact
//Email - nair.svnit@gmail.com
//linkedin - https://www.linkedin.com/in/karthiknair18/
//angellist - https://angel.co/u/karthik-nair-9
