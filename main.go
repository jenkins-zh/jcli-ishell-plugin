package main

import (
	inner "github.com/jenkins-zh/jcli-ishell-plugin/cmd"
	"os"
)

func main() {
	cmd := inner.NewJobCmd(os.Args[1:])
	cmd.Run()
}
