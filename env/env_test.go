package env

import (
	"os"
	"testing"
)

func TestInit(t *testing.T) {
	// Create a temporary .env file
	content := []byte("TEST_KEY=test_value")
	tmpfile, err := os.CreateTemp("", "test.env")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name()) // clean up

	if _, err := tmpfile.Write(content); err != nil {
		t.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		t.Fatal(err)
	}

	// Test Init with the temp file
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Init panicked: %v", r)
		}
	}()
	Init(tmpfile.Name())

	// Verify the env var is set
	if val := os.Getenv("TEST_KEY"); val != "test_value" {
		t.Errorf("Expected TEST_KEY=test_value, got %s", val)
	}
}

func TestGetEnvString(t *testing.T) {
	key := "MY_ENV_VAR"
	val := "my_value"
	os.Setenv(key, val)
	defer os.Unsetenv(key)

	if res := GetEnvString(key); res != val {
		t.Errorf("Expected %s, got %s", val, res)
	}
}

func TestInitPanic(t *testing.T) {
	// Test Init with non-existent file
	defer func() {
		if r := recover(); r == nil {
			t.Error("Init should have panicked with non-existent file")
		}
	}()
	Init("non_existent_file.env")
}
