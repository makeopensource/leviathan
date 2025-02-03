package main

import (
	"fmt"
	"github.com/makeopensource/leviathan/api"
	"github.com/makeopensource/leviathan/utils"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"net/http"
	"os"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	utils.InitConfig()

	mux := api.SetupEndpoints()

	port := "9221"
	srvAddr := fmt.Sprintf(":%s", port)
	log.Info().Msgf("Started server on %s", srvAddr)

	err := http.ListenAndServe(
		srvAddr,
		// Use h2c so we can serve HTTP/2 without TLS.
		h2c.NewHandler(mux, &http2.Server{}),
	)
	if err != nil {
		log.Fatal().Err(err).Msgf("Failed to start server on %s", srvAddr)
		return
	}
}
