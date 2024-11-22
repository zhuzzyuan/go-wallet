package config

import (
	"os"
)

// Running modes
const (
	DebugMode   = "debug"
	ReleaseMode = "release"
)

var modeName = DebugMode

func init() {
	SetMode(os.Getenv("MODE"))
}

// SetMode sets system running mode, e.g. config.DebugMode.
func SetMode(value string) {
	switch value {
	case DebugMode:
		modeName = DebugMode
	case ReleaseMode, "":
		modeName = ReleaseMode
	default:
		panic("system running mode unknown: " + value)
	}
}

// Mode returns current running mode.
func Mode() string {
	return modeName
}

// IsDebugMode tells if running in debug mode.
func IsDebugMode() bool {
	return modeName == DebugMode
}
