package main

import (
	"fmt"
	"github.com/punithsr27/arena/src/maths"
)

func main() {
	add := maths.Addition{}
	var value1, value2 float32
	_, err := fmt.Scanln("%f %f", &value1, &value2)
	if err != nil {
		add.FirstValue = value1
		add.SecondValue = value2
		fmt.Printf("Addition of two no's %f:", add.GetSum())
	} else {
		fmt.Println(err)
	}

}
