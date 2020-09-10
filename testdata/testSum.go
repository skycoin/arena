package testdata

import (
	"GolangDeveloper/src/sum"
	"fmt"
)

func TestTwoNumberSum() {
	var n1, n2 int
	fmt.Print("1st Number: ")
	_, n1Err := fmt.Scanf("%d", &n1)
	if n1Err != nil {
		fmt.Errorf(n1Err.Error())
	}
	fmt.Print("2st Number: ")
	_, n2Err := fmt.Scanf("%d", &n2)
	if n2Err != nil {
		fmt.Errorf(n2Err.Error())
	}
	fmt.Println(sum.FindSum().SumTwoNumbers(n1, n2))
}
