package arena

import (
	"fmt"
	"io"
)

const ErrorParsing = "Error when parsing input as a number: %v\n"

func ReadSumAndPrint(reader io.Reader, writer io.Writer, numbersToRead int) {
	var sum float64 = 0
	for i := 0; i < numbersToRead; {
		var number float64
		_, e := fmt.Fscan(reader, &number)
		if e != nil {
			_, e = fmt.Fprintf(writer, ErrorParsing, e.Error())
			if e != nil {
				fmt.Printf(ErrorParsing, e.Error())
			}
		} else {
			i++
			sum += number
		}
	}

	_, e := fmt.Fprint(writer, sum)
	if e != nil {
		fmt.Print(sum)
	}
}
