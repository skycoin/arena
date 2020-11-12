package cli

import (
	"github.com/gocodedrifter/arena/pkg/transaction"
)

// Command : command interface
type Command interface {
	Execute(x, y interface{}) (int, error)
}

// AddCommand : implement interface command
type AddCommand struct{}

// Execute : methode to implement
func (cmd *AddCommand) Execute(x, y interface{}) (int, error) {
	return transaction.AddTwoNumber(x, y)
}

// Exec : to execute
func Exec(cmd string, x, y interface{}) (int, error) {
	commands := map[string]Command{
		"add": &AddCommand{},
	}

	command := commands[cmd]
	return command.Execute(x, y)
}
