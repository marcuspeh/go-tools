package util

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEventManager_Subscribe(t *testing.T) {
	em := NewEventManager[string]()

	err := em.Subscribe("id1", func(data string) {})
	require.NoError(t, err)

	err = em.Subscribe("id1", func(data string) {})
	require.Error(t, err)
	require.Equal(t, "callback id already exists", err.Error())

	err = em.Subscribe("id2", func(data string) {})
	require.NoError(t, err)
}

func TestEventManager_Unsubscribe(t *testing.T) {
	em := NewEventManager[string]()

	em.Subscribe("id1", func(data string) {})
	em.Unsubscribe("id1")

	// Unsubscribe non-existent id should not panic
	require.NotPanics(t, func() {
		em.Unsubscribe("id1")
		em.Unsubscribe("non-existent")
	})
}

func TestEventManager_NotifyCallbacks(t *testing.T) {
	em := NewEventManager[string]()

	var received1 string
	var received2 string

	em.Subscribe("id1", func(data string) {
		received1 = data
	})
	em.Subscribe("id2", func(data string) {
		received2 = data
	})

	testData := "hello world"
	em.NotifyCallbacks("some-id", testData)

	require.Equal(t, testData, received1)
	require.Equal(t, testData, received2)
}

func TestEventManager_NotifyCallbacks_AfterUnsubscribe(t *testing.T) {
	em := NewEventManager[string]()

	var received1 string
	var received2 string

	em.Subscribe("id1", func(data string) {
		received1 = data
	})
	em.Subscribe("id2", func(data string) {
		received2 = data
	})

	em.Unsubscribe("id1")

	testData := "hello world"
	em.NotifyCallbacks("some-id", testData)

	require.Equal(t, "", received1)
	require.Equal(t, testData, received2)
}

func TestEventManager_ComplexType(t *testing.T) {
	type TestStruct struct {
		Value int
	}
	em := NewEventManager[*TestStruct]()

	var receivedValue int
	em.Subscribe("id1", func(data *TestStruct) {
		receivedValue = data.Value
	})

	em.NotifyCallbacks("id1", &TestStruct{Value: 42})
	require.Equal(t, 42, receivedValue)
}
