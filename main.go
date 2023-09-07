package main

import (
	"fmt"
	"grl/cmd"
)

func main() {
	fmt.Println("Hello GRL")
	app := cmd.NewCli()
	app.Run()
}
