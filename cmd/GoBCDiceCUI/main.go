package main
import (
	"fmt"
	"os"
	"github.com/urfave/cli/v2"
)

func main() {
	os.Exit(run())
}

func msg(err error) int {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", os.Args[0], err)
		return 1
	}
	return 0
}

func appRun(c *cli.Context)  error {
	args := c.Args()
	if !args.Present() {
		cli.ShowAppHelp(c)
		return nil
	}
	return nil
}

func run() int {
	app := cli.NewApp()
	app.Name = "GoBCDice"
	app.Usage = "BCDice Compatible Command"
	app.Action = appRun
	return msg(app.Run(os.Args))
}
