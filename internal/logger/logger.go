package logger

import (
	"log/slog"
	"os"
)

var Logger *slog.Logger

func NewLogger() {
	Logger = slog.New(slog.NewTextHandler(os.Stdout, nil))
}
