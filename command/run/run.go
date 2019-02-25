package run

import (
	"errors"
	"fmt"

	"github.com/inconshreveable/log15"
	"github.com/mrcaelumn/arena/services"
	"github.com/urfave/cli"
)

var Command = cli.Command{
	Name:  "run",
	Usage: "Run the service",
	Flags: []cli.Flag{
		cli.IntSliceFlag{
			Name:  "input, i",
			Usage: "Input Ex: 1 ",
		},
	},
	Action: func(c *cli.Context) error {
		log := log15.New("module", "arenaSkycoin")
		if len(c.IntSlice("input")) == 0 {
			// fmt.Println("input is empty")
			return errors.New("input is empty")
		}
		log.Debug("Input: ", c.IntSlice("input"))
		appserv := services.New(c.IntSlice("input"))
		result, err := appserv.SumsUp()
		if err != nil {
			fmt.Println(err)
			return err
		}
		fmt.Println(result)
		return nil
	},
}
