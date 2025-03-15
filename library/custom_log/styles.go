package custom_log

import (
	"fmt"
	"os"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
)

type CustomLogger struct {
	*log.Logger
}

const (
	SuccessLevel log.Level = 2
	timeFormat string = "2006/01/02 15:04:05"
)

var (
	Logger *CustomLogger = nil
)

func (l *CustomLogger) Success(msg interface{}, keyvals ...interface{}) {
	l.Log(SuccessLevel, msg, keyvals...)
}

func (l *CustomLogger) Successf(format string, args ...interface{}) {
	l.Log(SuccessLevel, fmt.Sprintf(format, args...))
}

func PrepareLogger() {
	styles := log.DefaultStyles()
	styles.Levels[SuccessLevel] = lipgloss.NewStyle().
		SetString("SUCCESS").
		Foreground(lipgloss.Color("#1eff00")).Bold(true)

	styles.Keys["success"] = lipgloss.NewStyle().Foreground(lipgloss.Color("#597445"))
	styles.Keys["key"] = lipgloss.NewStyle().Foreground(lipgloss.Color("#597445"))

	Logger = &CustomLogger{log.New(os.Stderr)}
	Logger.SetLevel(SuccessLevel)
	Logger.SetReportTimestamp(true)
	Logger.SetStyles(styles)
}
