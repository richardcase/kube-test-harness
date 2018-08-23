package logger

import "github.com/dlespiau/kube-test-harness/testingiface"

// TestLogger is a logger using testing.T.Log for its output.
type TestLogger struct {
	baseLogger
}

var _ Logger = &TestLogger{}

// ForTest implements Logger.
func (l *TestLogger) ForTest(t testingiface.TestingT) Logger {
	return &TestLogger{
		baseLogger: baseLogger{
			level: l.level,
			t:     t,
		},
	}
}

// Log implements Logger.
func (l *TestLogger) Log(level LogLevel, msg string) {
	if !l.shouldLog(level) {
		return
	}
	if l.t != nil {
		if h, ok := l.t.(testingiface.HelperT); ok {
			h.Helper()
		}
		l.t.Log(msg)
		return
	}
	pl := PrintfLogger{}
	pl.Log(level, msg)
}

// Logf implements Logger.
func (l *TestLogger) Logf(level LogLevel, f string, args ...interface{}) {
	if !l.shouldLog(level) {
		return
	}
	if l.t != nil {
		if h, ok := l.t.(testingiface.HelperT); ok {
			h.Helper()
		}
		l.t.Logf(f, args...)
		return
	}
	pl := PrintfLogger{}
	pl.Logf(level, f, args...)
}
