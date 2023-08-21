package log

import (
	"github.com/rs/zerolog"

	"github.com/ggymm/gopkg/log"
)

var isInit = false
var logger zerolog.Logger

func Init(filename string) {
	isInit = true
	logger = log.InitCustom(filename)
}
