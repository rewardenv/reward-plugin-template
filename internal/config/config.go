package config

import (
	"fmt"
	"path"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/hashicorp/go-version"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/afero"
	"github.com/spf13/viper"
)

var (
	// FS is the implementation of Afero Filesystem. It's a filesystem wrapper and used for testing.
	FS = &afero.Afero{Fs: afero.NewOsFs()}
)

type Config struct {
	*viper.Viper
	appName string
}

func New(name, ver string) *Config {
	c := &Config{
		Viper:   viper.GetViper(),
		appName: name,
	}

	c.SetDefault(fmt.Sprintf("%s_version", name), version.Must(version.NewVersion(ver)).String())

	return c
}

func (c *Config) Init() *Config {
	c.AddConfigPath(".")

	cfg := c.GetString(fmt.Sprintf("%s_config_file", c.appName))
	if cfg != "" {
		c.AddConfigPath(filepath.Dir(cfg))
		c.SetConfigName(filepath.Base(cfg))
		c.SetConfigType("yaml")
	}
	c.AutomaticEnv()

	if err := c.ReadInConfig(); err != nil {
		log.Debugf("%v", err)
	}

	c.AddConfigPath(".")
	c.SetConfigName(".env")
	c.SetConfigType("dotenv")

	if err := c.MergeInConfig(); err != nil {
		log.Debugf("%v", err)
	}

	c.SetLogging()

	return c
}

// SetLogging sets the logging level based on the command line flags and environment variables.
func (c *Config) SetLogging() {
	switch {
	case c.GetString("log_level") == "trace":
		log.SetLevel(log.TraceLevel)
		log.SetReportCaller(true)
	case c.IsDebug(), c.GetString("log_level") == "debug":
		log.SetLevel(log.DebugLevel)
		log.SetReportCaller(true)
	case c.GetString("log_level") == "info":
		log.SetLevel(log.InfoLevel)
	case c.GetString("log_level") == "warning":
		log.SetLevel(log.WarnLevel)
	default:
		log.SetLevel(log.ErrorLevel)
	}

	log.SetFormatter(
		&log.TextFormatter{
			DisableColors:          c.GetBool("disable_colors"),
			ForceColors:            true,
			DisableLevelTruncation: true,
			FullTimestamp:          true,
			DisableTimestamp:       !c.GetBool("debug"),
			QuoteEmptyFields:       true,
			CallerPrettyfier: func(f *runtime.Frame) (string, string) {
				filename := strings.ReplaceAll(path.Base(f.File), "reward/", "")

				return fmt.Sprintf("%s()", f.Function), fmt.Sprintf(" %s:%d", filename, f.Line)
			},
		},
	)
}

// IsDebug returns true if debug mode is set.
func (c *Config) IsDebug() bool {
	return c.GetBool("debug")
}

func (c *Config) AppName() string {
	return c.appName
}

func (c *Config) AppVersion() string {
	return c.GetString(fmt.Sprintf("%s_version", c.appName))
}
