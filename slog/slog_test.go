package slog

import (

	"testing"
)

func TestInitLog(t *testing.T) {
	InitLog("Test")
}

func TestDebug(t *testing.T) {
	InitLog("Test")
	Debug("dfafasdfdsf")
}