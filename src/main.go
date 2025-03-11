package main

import (
	"github.com/makeopensource/leviathan/api"
	"github.com/makeopensource/leviathan/common"
	"github.com/rs/zerolog/log"
)

func main() {
	common.PrintInfo()
	log.Logger = common.ConsoleLogger()
	common.InitConfig()
	api.StartGrpcServer()
}
