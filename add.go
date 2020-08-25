package main;


import (
	"fmt"
	"strconv"
	"errors"
	"os"
)


func add(a interface{}, b  interface{}) (interface{}, error){

	aNum, err1 := strconv.Atoi(a.(string));
	bNum, err2 :=strconv.Atoi(b.(string));

	if err1!=nil && err2!=nil {
		return nil, errors.New("Both arguments are not integers")
	}

	if err1!=nil {
		return nil, errors.New("First argument is not an integer")
	}

	if err2!=nil {
		return nil, errors.New("Second argument is not an integer")
	}

	return (aNum+bNum) , nil
}


func main() {

	var arg1 = os.Args[1]
	var arg2 = os.Args[2]

	result, err:= add(arg1, arg2)

	if err!=nil {
		fmt.Println("Error : " , err);
	}

	fmt.Println(result)

}