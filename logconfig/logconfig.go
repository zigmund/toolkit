package logconfig

import (
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
)

type (
	Options struct {
		Level        string `yaml:"level"`
		Format       string `yaml:"format"`
		Timestamp    bool   `yaml:"timestamp"`
		ReportCaller bool   `yaml:"reportCaller"`
	}
)

const (
	FormatText = "text"
	FormatJSON = "json"
)

// Configure - configures logrus using Options
func Configure(options Options) (err error) {
	lever, err := logrus.ParseLevel(options.Level)
	if err != nil {
		return err
	}

	logrus.SetLevel(lever)
	logrus.SetReportCaller(options.ReportCaller)

	switch options.Format {
	case FormatJSON:
		logrus.SetFormatter(&logrus.JSONFormatter{DisableTimestamp: !options.Timestamp})
	case FormatText:
		logrus.SetFormatter(&logrus.TextFormatter{DisableTimestamp: !options.Timestamp, FullTimestamp: true})
	default:
		return errors.New(fmt.Sprintf("unknown logger format: %s", options.Format))
	}

	return nil
}

// MustConfigure - configures logrus using Options or throws fatal on error
func MustConfigure(options Options) {
	if err := Configure(options); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
