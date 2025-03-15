package custom_log

import (
	"fmt"
	"os"
	"sync"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
)

type CustomLogger struct {
	*log.Logger
}

const (
	SuccessLevel log.Level = 2
)

var (
	L                 *CustomLogger = nil
	defaultLoggerOnce sync.Once
)

// Default returns the default L. The default L comes with timestamp enabled.
func Default() *CustomLogger {
	defaultLoggerOnce.Do(
		func() {
			if L != nil {
				// already set via SetDefault.
				return
			}
			styles := log.DefaultStyles()
			styles.Levels[SuccessLevel] = lipgloss.NewStyle().
				SetString("SUCCESS").
				Foreground(lipgloss.Color("#1eff00")).Bold(true)

			styles.Keys["success"] = lipgloss.NewStyle().Foreground(lipgloss.Color("#597445"))
			styles.Keys["key"] = lipgloss.NewStyle().Foreground(lipgloss.Color("#597445"))

			L = &CustomLogger{log.New(os.Stderr)}
			L.SetLevel(log.DebugLevel)
			L.SetReportTimestamp(true)
			L.SetStyles(styles)
		},
	)
	return L
}

func (l *CustomLogger) Success(msg interface{}, keyvals ...interface{}) {
	Default().Log(SuccessLevel, msg, keyvals...)
}

func (l *CustomLogger) Successf(format string, args ...interface{}) {
	Default().Log(SuccessLevel, fmt.Sprintf(format, args...))
}
