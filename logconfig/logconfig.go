package logconfig

import (
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/writer"
	"io"
	"os"
)

type (
	Options struct {
		Level        string `yaml:"level"`
		Format       string `yaml:"format"`
		Timestamp    bool   `yaml:"timestamp"`
		ReportCaller bool   `yaml:"reportCaller"`
		SplitOutput  bool   `yaml:"splitOutput"`
	}
)

const (
	FormatText = "text"
	FormatJSON = "json"
)

// Configure - configures logrus using Options
func Configure(options Options) (err error) {
	level, err := logrus.ParseLevel(options.Level)
	if err != nil {
		return err
	}

	logrus.SetLevel(level)
	logrus.SetReportCaller(options.ReportCaller)

	switch options.Format {
	case FormatJSON:
		logrus.SetFormatter(&logrus.JSONFormatter{DisableTimestamp: !options.Timestamp})
	case FormatText:
		logrus.SetFormatter(&logrus.TextFormatter{DisableTimestamp: !options.Timestamp, FullTimestamp: true})
	default:
		return errors.New(fmt.Sprintf("unknown logger format: %s", options.Format))
	}

	// https://github.com/sirupsen/logrus/tree/f104497f2b2129ab888fd274891f3a278756bcde/hooks/writer
	if options.SplitOutput {
		logrus.SetOutput(io.Discard)
		logrus.AddHook(&writer.Hook{
			Writer: os.Stderr,
			LogLevels: []logrus.Level{
				logrus.PanicLevel,
				logrus.FatalLevel,
				logrus.ErrorLevel,
				logrus.WarnLevel,
			},
		})
		logrus.AddHook(&writer.Hook{
			Writer: os.Stdout,
			LogLevels: []logrus.Level{
				logrus.InfoLevel,
				logrus.DebugLevel,
				logrus.TraceLevel,
			},
		})
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
