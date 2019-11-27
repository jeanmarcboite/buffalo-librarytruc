package net

import (
	"os"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/rawbytes"

	"github.com/jeanmarcboite/bookins/pkg/books/online/assets"
	"github.com/sirupsen/logrus"
)

// Koanf -- Global koanf instance. Use . as the key path delimiter. This can be / or anything.
var Koanf = koanf.New(".")

// Logger -- metadata logger
var Logger = logrus.New()

func init() {
	Logger.Out = os.Stdout
	Logger.SetLevel(logrus.DebugLevel)
	conf, err := assets.Config.Find("urls.yaml")
	if err == nil {
		Koanf.Load(rawbytes.Provider(conf), yaml.Parser())
	}
}
