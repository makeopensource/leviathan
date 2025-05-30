package main

import (
	"github.com/makeopensource/leviathan/cmd"
	"github.com/makeopensource/leviathan/internal/info"
)

func main() {
	info.PrintInfo()
	cmd.Setup()
	cmd.StartServerWithAddr()
}
