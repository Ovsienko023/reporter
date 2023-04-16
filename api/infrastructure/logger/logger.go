package logger

import (
	"github.com/Ovsienko023/reporter/infrastructure/configuration"
	"go.uber.org/zap"
	"gopkg.in/natefinch/lumberjack.v2"
	"net/url"
	"os"
	"path/filepath"
	"runtime"
)

func New(cfg *configuration.Api) (*zap.Logger, error) {
	logLevel := `debug`

	logDir := `/opt/trueconf`
	logName := "api" + ".log"

	_ = os.MkdirAll(logDir, 0777)

	var level zap.AtomicLevel

	switch logLevel {
	case `debug`:
		level = zap.NewAtomicLevelAt(zap.DebugLevel)
	case `warning`:
		level = zap.NewAtomicLevelAt(zap.WarnLevel)
	case `error`:
		level = zap.NewAtomicLevelAt(zap.ErrorLevel)
	case `panic`:
		level = zap.NewAtomicLevelAt(zap.PanicLevel)
	case `fatal`:
		level = zap.NewAtomicLevelAt(zap.FatalLevel)
	default:
		level = zap.NewAtomicLevelAt(zap.InfoLevel)
	}

	var conf zap.Config

	if err := zap.RegisterSink("rotate", func(u *url.URL) (zap.Sink, error) {
		filename := u.Path

		if runtime.GOOS == "windows" {
			filename = u.Opaque
		}

		return &rotateLogger{
			Logger: &lumberjack.Logger{
				Filename:   filename,
				MaxSize:    30,
				MaxBackups: 10,
				MaxAge:     0,
			},
		}, nil
	}); err != nil {
		return nil, err
	}

	path := "rotate:" + filepath.Join(logDir, logName)

	conf = zap.NewDevelopmentConfig()

	conf.Level = level
	conf.DisableStacktrace = true
	conf.OutputPaths = []string{"stdout", path}
	conf.ErrorOutputPaths = []string{"stdout", path}

	return conf.Build()
}

type rotateLogger struct{ *lumberjack.Logger }

func (r *rotateLogger) Sync() error { return r.Close() }
