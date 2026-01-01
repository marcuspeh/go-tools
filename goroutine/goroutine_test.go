package goroutine

import (
	"context"
	"errors"
	"testing"
)

func TestErrGroup_Success(t *testing.T) {
	g := NewErrGroup()
	ctx := context.Background()

	runCount := 0
	g.Run(ctx, func() error {
		runCount++
		return nil
	})

	if err := g.Wait(); err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if runCount != 1 {
		t.Errorf("Expected runCount 1, got %d", runCount)
	}
}

func TestErrGroup_Error(t *testing.T) {
	g := NewErrGroup()
	ctx := context.Background()
	expectedErr := errors.New("test error")

	g.Run(ctx, func() error {
		return expectedErr
	})

	if err := g.Wait(); err != expectedErr {
		t.Errorf("Expected error %v, got %v", expectedErr, err)
	}
}

func TestErrGroup_Panic(t *testing.T) {
	g := NewErrGroup()
	ctx := context.Background()

	g.Run(ctx, func() error {
		panic("test panic")
	})

	err := g.Wait()
	if err == nil {
		t.Fatal("Expected error from panic, got nil")
	}

	if err.Error() != "panic occured test panic" {
		t.Errorf("Expected panic error message, got %v", err)
	}
}
