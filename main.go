package main

//go:generate fileb0x fileb0x.yml

import "github.com/boundedinfinity/userdotd/cmd"

func main() {
	cmd.Execute()
}
