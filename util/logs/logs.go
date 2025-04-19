package logs

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger is the default global logger, it is a no-op dumb logger before initialization, and does nothing.
var Logger = zap.NewNop().Sugar()

// Error creates an error field from the given error.
func Error(err error) zapcore.Field {
	if err == nil {
		return zap.Skip()
	}
	return zap.Error(err)
}
