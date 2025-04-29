package file_manager

import (
	"github.com/rs/zerolog/log"
	"io"
)

// closeWithLog closes an io.closer
// prints a warning log if an error occurs
func closeWithLog(closer io.Closer) {
	err := closer.Close()
	if err != nil {
		log.Warn().Err(err).Msg("Error occurred while closing interface")
	}
}
