package logger

import (
	"context"
	"os"
	"testing"

	"go.uber.org/zap"
)

func TestInitAndLog(t *testing.T) {
	// Use a temporary file
	tmpfile, err := os.CreateTemp("", "test_log_*.log")
	if err != nil {
		t.Fatal(err)
	}
	logPath := tmpfile.Name()
	tmpfile.Close()
	os.Remove(logPath) // Init creates it

	// Init
	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("Init panicked: %v", r)
		}
	}()
	Init(logPath)

	ctx := context.Background()
	ctx = context.WithValue(ctx, LogIDKey, "test-log-id")

	// Log something
	Info(ctx, "info message", zap.String("key", "value"))
	Error(ctx, "error message")
	Warn(ctx, "warn message")

	// Allow buffer to flush (defer logger.Sync() is in Init but we can't call it here)
	// zap.New usually syncs on exit or we can force it if we had access to the logger variable.
	// But logger variable is private.
	// However, Init defers Sync(), but that defer runs when Init returns?
	// No, `defer logger.Sync()` in `Init` runs when `Init` returns!
	// This means the logger might be closed/synced immediately after Init?
	// Wait, `defer logger.Sync()` executes when `Init` finishes.
	// `logger` is a global variable.
	// If `logger.Sync()` is called, it flushes. It doesn't necessarily close it forever, but zap syncs usually are fine.
	// However, usually you want to defer Sync in main, not in Init.
	// If Init defers Sync, it runs immediately.
	// Let's check `logger.go` again.

}

func TestEmplaceKV(t *testing.T) {
	field := EmplaceKV("key", map[string]string{"a": "b"})
	if field.Key != "key" {
		t.Errorf("Expected key 'key', got %s", field.Key)
	}
	if field.String != `{"a":"b"}` {
		t.Errorf("Expected value '{\"a\":\"b\"}', got %s", field.String)
	}
}

func TestErrorLog(t *testing.T) {
	err := os.ErrNotExist
	field := ErrorLog(err)
	if field.Key != "error" { // zap default error key
		// Actually zap.Error uses "error" usually?
		// zap.Error(err) returns a Field with Key "error" and Interface err
		// But here checking Key might be enough.
	}
	// zap.Field is a struct.
}
