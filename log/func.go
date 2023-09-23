package log

import (
	"github.com/rs/zerolog"

	"github.com/ggymm/gopkg/logger"
)

func autoInit() {
	if !isInit {
		// 手动初始化
		logger.Init()
	}
}

func Trace() *zerolog.Event {
	autoInit()
	return log.Trace()
}

func Debug() *zerolog.Event {
	autoInit()
	return log.Debug()
}

func Info() *zerolog.Event {
	autoInit()
	return log.Info()
}

func Warn() *zerolog.Event {
	autoInit()
	return log.Warn()
}

func Error() *zerolog.Event {
	autoInit()
	return log.Error().Stack()
}

func Fatal() *zerolog.Event {
	autoInit()
	return log.Fatal()
}

func Panic() *zerolog.Event {
	autoInit()
	return log.Panic().Stack()
}
