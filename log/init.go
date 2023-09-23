package log

import (
	"github.com/rs/zerolog"

	"github.com/ggymm/gopkg/logger"
)

var isInit = false
var log zerolog.Logger

func Init(filename string) {
	isInit = true
	log = logger.Init(filename)
}
