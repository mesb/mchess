// store information about the configuration of the mchess
package config // import "github.com/mesb/mchess/config"

import (
	"io"
)

type Config struct {
	prompt   string
	output   io.Writer
	errorput io.Writer
}
