package main

import (
	"fmt"
	"os"

	"github.com/thejunghare/domain/cmd"
)

func main() {
	checkerror(cmd.RootCmd.Execute())
}

func checkerror(err error) {
	if err != nil {
		fmt.Print("Something went wrong", err.Error())
		os.Exit(1)
	}
}
