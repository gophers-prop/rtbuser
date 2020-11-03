package main

import (
	"os"
	"rtbuser/cmd"
)


func main() {
	cmd.Commands(os.Args[1:])
}