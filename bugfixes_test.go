package bugfixes

import "testing"

// LOG
func TestLog(t *testing.T) {
	Log("Log Test")
}

// WARN
func TestInfo(t *testing.T) {
	Info("Info Test")
}

func TestWarn(t *testing.T) {
	Warn("Warn Test")
}

// ERROR
func TestError(t *testing.T) {
	Error("Error Test")
}

func TestFatal(t *testing.T) {
	Fatal("Fatal Test")
}