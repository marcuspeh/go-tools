package ctx

import (
	"strings"
	"testing"
	"time"

	"github.com/marcuspeh/go-tools/logger"
)

func TestGetCtx(t *testing.T) {
	postfix := "test-postfix"
	ctx, cancel := GetCtx(postfix)
	defer cancel()

	if ctx == nil {
		t.Fatal("Expected context, got nil")
	}

	logID := ctx.Value(logger.LogIDKey)
	if logID == nil {
		t.Fatal("Expected LogID in context, got nil")
	}

	logIDStr, ok := logID.(string)
	if !ok {
		t.Fatal("Expected LogID to be string")
	}

	if !strings.Contains(logIDStr, postfix) {
		t.Errorf("Expected LogID %s to contain postfix %s", logIDStr, postfix)
	}

	// Verify format somewhat (Date_Time_Postfix)
	parts := strings.Split(logIDStr, "_")
	if len(parts) < 3 {
		t.Errorf("Expected at least 3 parts in LogID, got %d", len(parts))
	}

	// Verify date format
	_, err := time.Parse(time.DateOnly, parts[0])
	if err != nil {
		t.Errorf("Invalid date format in LogID: %v", err)
	}

	// Verify time format
	_, err = time.Parse(time.TimeOnly, parts[1])
	if err != nil {
		t.Errorf("Invalid time format in LogID: %v", err)
	}

	// Verify cancellation
	select {
	case <-ctx.Done():
		t.Fatal("Context should not be done yet")
	default:
	}

	cancel()

	select {
	case <-ctx.Done():
	case <-time.After(1 * time.Second):
		t.Fatal("Context should be done after cancel")
	}
}
