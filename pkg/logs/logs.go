package logs

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var zp zapInstance

type zapInstance struct {
	Log *zap.Logger
}

func InitZapLog() {
	conf := zap.NewProductionConfig()
	conf.DisableStacktrace = true
	conf.EncoderConfig.TimeKey = "timestamp"
	conf.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	zaplog, err := conf.Build(zap.AddCallerSkip(1))
	if err != nil {
		zaplog.Error(err.Error())
	}

	zp = zapInstance{
		Log: zaplog,
	}
}

func Info(message string, fields ...zap.Field) {
	zp.Log.Info(message, fields...)
}

func Error(message interface{}, fields ...zap.Field) {
	switch val := message.(type) {
	case error:
		zp.Log.Error(val.Error(), fields...)
	case string:
		zp.Log.Error(val, fields...)
	default:
		zp.Log.Error("unknown error", fields...)
	}
}
