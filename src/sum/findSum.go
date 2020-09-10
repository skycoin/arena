package sum

import "fmt"

type findSum struct{}

type FindSumIF interface {
	SumTwoNumbers(n1 int, n2 int) (int, error)
}

func FindSum() FindSumIF {
	return &findSum{}
}

func (self *findSum) SumTwoNumbers(n1 int, n2 int) (int, error) {
	if n1 == 0 {
		fmt.Print("You havent provided a Valid 1st Number: ")
		_, n1Err := fmt.Scanf("%d", &n1)
		if n1Err != nil {
			fmt.Errorf(n1Err.Error())
			return 0, n1Err
		}
	}
	if n2 == 0 {
		fmt.Print("You havent provided a Valid  2st Number: ")
		_, n2Err := fmt.Scanf("%d", &n2)
		if n2Err != nil {
			fmt.Errorf(n2Err.Error())
			return 0, n2Err
		}
	}
	return n1 + n2, nil
}
