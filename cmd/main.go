package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/gocodedrifter/arena/src/cli"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	params := scanner.Text()
	if words := strings.Fields(params); len(words) > 1 {
		if result, err := cli.Exec("add", words[0], words[1]); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(result)
		}
	}

}
