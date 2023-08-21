package log

import (
	"github.com/rs/zerolog"

	"github.com/ggymm/gopkg/log"
)

func autoInit() {
	if !isInit {
		// 手动初始化
		log.InitCustom()
	}
}

func Trace() *zerolog.Event {
	autoInit()
	return logger.Trace()
}

func Debug() *zerolog.Event {
	autoInit()
	return logger.Debug()
}

func Info() *zerolog.Event {
	autoInit()
	return logger.Info()
}

func Warn() *zerolog.Event {
	autoInit()
	return logger.Warn()
}

func Error() *zerolog.Event {
	autoInit()
	return logger.Error().Stack()
}

func Fatal() *zerolog.Event {
	autoInit()
	return logger.Fatal()
}

func Panic() *zerolog.Event {
	autoInit()
	return logger.Panic().Stack()
}
