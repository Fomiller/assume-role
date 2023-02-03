package main

import (
	"github.com/Fomiller/assume-role/cmd"
)

type assumeConfig struct {
	account string
	role    string
}

func main() {
	cmd.Execute()
}
