package internal

import (
	"fmt"
	"io"
	"log/slog"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func SetupLogger() {
	lvlStr := viper.GetString("logLevel")
	lvl, err := getLogLevel(lvlStr)
	cobra.CheckErr(err)
	logPath := viper.GetString("logPath")
	logFile, err := openLogFile(logPath)
	cobra.CheckErr(err)
	handler := slog.NewTextHandler(logFile, &slog.HandlerOptions{Level: lvl})
	logger := slog.New(handler)
	slog.SetDefault(logger)
}

func getLogLevel(lvl string) (slog.Level, error) {
	switch lvl {
	case "error":
		return slog.LevelError, nil
	case "warn":
		return slog.LevelWarn, nil
	case "info":
		return slog.LevelInfo, nil
	case "debug":
		return slog.LevelDebug, nil
	}

	return -100, fmt.Errorf("couldn't determine log level")
}

func openLogFile(path string) (io.Writer, error) {
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}
	return file, nil
}
