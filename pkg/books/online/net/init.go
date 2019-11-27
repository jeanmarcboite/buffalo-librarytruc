package net

import (
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/rawbytes"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/jeanmarcboite/bookins/pkg/books/online/assets"
)

// Koanf -- Global koanf instance. Use . as the key path delimiter. This can be / or anything.
var Koanf = koanf.New(".")

// Logger -- metadata logger
var Logger *zap.SugaredLogger

func init() {
	loggerConfig := zap.NewDevelopmentConfig()
	loggerConfig.Level = zap.NewAtomicLevelAt(zapcore.DebugLevel)
	logger, _ := loggerConfig.Build()
	logger.Sugar()
	Logger = logger.Sugar()
	// Logger.SetLevel(logrus.DebugLevel)
	conf, err := assets.Config.Find("urls.yaml")
	if err == nil {
		Koanf.Load(rawbytes.Provider(conf), yaml.Parser())
	}
}
