package tools

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestThreadSafeMap_Get(t *testing.T) {
	tests := []struct {
		name        string
		key         string
		before      func(m ThreadSafeMap)
		threads     int
		expectedVal interface{}
		expectedOk  bool
	}{
		{
			name: "get",
			key:  "key",
			before: func(m ThreadSafeMap) {
				m.(*ThreadSafeMapImpl).internalMap["key"] = "value"
			},
			threads:     1,
			expectedVal: "value",
			expectedOk:  true,
		},
		{
			name:        "get not found",
			key:         "key",
			before:      func(m ThreadSafeMap) {},
			threads:     1,
			expectedVal: nil,
			expectedOk:  false,
		},
		{
			name: "get concurrent",
			key:  "key",
			before: func(m ThreadSafeMap) {
				m.(*ThreadSafeMapImpl).internalMap["key"] = "value"
			},
			threads:     100,
			expectedVal: "value",
			expectedOk:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := NewThreadSafeMap()
			tt.before(m)

			wg := sync.WaitGroup{}
			wg.Add(tt.threads)
			for i := 0; i < tt.threads; i++ {
				go func() {
					defer wg.Done()
					val, ok := m.Get(tt.key)
					require.Equal(t, tt.expectedVal, val)
					require.Equal(t, tt.expectedOk, ok)
				}()
			}
			wg.Wait()
		})
	}
}

func TestThreadSafeMap_Set(t *testing.T) {
	tests := []struct {
		name        string
		key         string
		val         interface{}
		before      func(m ThreadSafeMap)
		threads     int
		expectedVal interface{}
		expectedOk  bool
	}{
		{
			name:        "set",
			key:         "key",
			val:         "value",
			threads:     1,
			before:      func(m ThreadSafeMap) {},
			expectedVal: "value",
			expectedOk:  true,
		},
		{
			name:        "set concurrent",
			key:         "key",
			val:         "value",
			before:      func(m ThreadSafeMap) {},
			threads:     100,
			expectedVal: "value",
			expectedOk:  true,
		},
		{
			name: "set overwrite",
			key:  "key",
			val:  "value",
			before: func(m ThreadSafeMap) {
				m.(*ThreadSafeMapImpl).internalMap["key"] = "value"
			},
			threads:     1,
			expectedVal: "value",
			expectedOk:  true,
		},
		{
			name: "set overwrite concurrent",
			key:  "key",
			val:  "value",
			before: func(m ThreadSafeMap) {
				m.(*ThreadSafeMapImpl).internalMap["key"] = "value"
			},
			threads:     100,
			expectedVal: "value",
			expectedOk:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := NewThreadSafeMap()
			tt.before(m)

			wg := sync.WaitGroup{}
			wg.Add(tt.threads)
			for i := 0; i < tt.threads; i++ {
				go func() {
					defer wg.Done()
					m.Set(tt.key, tt.val)
				}()
			}

			wg.Wait()
			val, ok := m.(*ThreadSafeMapImpl).internalMap[tt.key]
			require.Equal(t, tt.expectedVal, val)
			require.Equal(t, tt.expectedOk, ok)
		})
	}
}

func TestThreadSafeMap_Update(t *testing.T) {
	tests := []struct {
		name        string
		key         string
		val         interface{}
		before      func(m ThreadSafeMap)
		threads     int
		fn          func(oldVal, newVal interface{}) interface{}
		expectedVal interface{}
		expectedOk  bool
	}{
		{
			name:    "update",
			key:     "key",
			val:     "value",
			before:  func(m ThreadSafeMap) {},
			threads: 1,
			fn: func(oldVal, newVal interface{}) interface{} {
				return newVal
			},
			expectedVal: "value",
			expectedOk:  true,
		},
		{
			name:    "update concurrent",
			key:     "key",
			val:     "value",
			before:  func(m ThreadSafeMap) {},
			threads: 100,
			fn: func(oldVal, newVal interface{}) interface{} {
				return newVal
			},
			expectedVal: "value",
			expectedOk:  true,
		},
		{
			name: "update overwrite",
			key:  "key",
			val:  "value",
			before: func(m ThreadSafeMap) {
				m.Set("key", "value")
			},
			threads: 1,
			fn: func(oldVal, newVal interface{}) interface{} {
				return newVal
			},
			expectedVal: "value",
			expectedOk:  true,
		},
		{
			name: "update overwrite concurrent",
			key:  "key",
			val:  "value",
			before: func(m ThreadSafeMap) {
				m.Set("key", "value")
			},
			threads: 100,
			fn: func(oldVal, newVal interface{}) interface{} {
				return newVal
			},
			expectedVal: "value",
			expectedOk:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := NewThreadSafeMap()
			tt.before(m)

			wg := sync.WaitGroup{}
			wg.Add(tt.threads)
			for i := 0; i < tt.threads; i++ {
				go func() {
					defer wg.Done()
					m.Update(tt.key, tt.val, tt.fn)
				}()
			}
			wg.Wait()
			val, ok := m.(*ThreadSafeMapImpl).internalMap[tt.key]
			require.Equal(t, tt.expectedVal, val)
			require.Equal(t, tt.expectedOk, ok)
		})
	}
}

func TestThreadSafeMap_AddListKey(t *testing.T) {
	expectedValForConcurrent := []interface{}{}
	for i := 0; i < 100; i++ {
		expectedValForConcurrent = append(expectedValForConcurrent, "value")
	}

	tests := []struct {
		name        string
		key         string
		val         interface{}
		before      func(m ThreadSafeMap)
		threads     int
		expectedVal []interface{}
		expectedOk  bool
	}{
		{
			name:        "add list key",
			key:         "key",
			val:         "value",
			before:      func(m ThreadSafeMap) {},
			threads:     1,
			expectedVal: []interface{}{"value"},
			expectedOk:  true,
		},
		{
			name:        "add list key concurrent",
			key:         "key",
			val:         "value",
			before:      func(m ThreadSafeMap) {},
			threads:     100,
			expectedVal: expectedValForConcurrent,
			expectedOk:  true,
		},
		{
			name: "add list key exisiting",
			key:  "key",
			val:  "value",
			before: func(m ThreadSafeMap) {
				m.(*ThreadSafeMapImpl).internalMap["key"] = []interface{}{"value"}
			},
			threads:     1,
			expectedVal: []interface{}{"value", "value"},
			expectedOk:  true,
		},
		{
			name: "add list key exisiting concurrent",
			key:  "key",
			val:  "value",
			before: func(m ThreadSafeMap) {
				m.(*ThreadSafeMapImpl).internalMap["key"] = []interface{}{"value"}
			},
			threads:     99,
			expectedVal: expectedValForConcurrent,
			expectedOk:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := NewThreadSafeMap()
			tt.before(m)

			wg := sync.WaitGroup{}
			wg.Add(tt.threads)
			for i := 0; i < tt.threads; i++ {
				go func() {
					defer wg.Done()
					m.AddListItem(tt.key, tt.val)
				}()
			}
			wg.Wait()
			val, ok := m.(*ThreadSafeMapImpl).internalMap[tt.key]
			require.Equal(t, tt.expectedVal, val)
			require.Equal(t, tt.expectedOk, ok)
		})
	}
}

func TestThreadSafeMap_Del(t *testing.T) {
	tests := []struct {
		name        string
		key         string
		before      func(m ThreadSafeMap)
		threads     int
		expectedVal interface{}
		expectedOk  bool
	}{
		{
			name: "delete",
			key:  "key",
			before: func(m ThreadSafeMap) {
				m.(*ThreadSafeMapImpl).internalMap["key"] = "value"
			},
			threads:     1,
			expectedVal: nil,
			expectedOk:  false,
		},
		{
			name: "delete concurrent",
			key:  "key",
			before: func(m ThreadSafeMap) {
				m.(*ThreadSafeMapImpl).internalMap["key"] = "value"
			},
			threads:     100,
			expectedVal: nil,
			expectedOk:  false,
		},
		{
			name:        "delete not found",
			key:         "key",
			before:      func(m ThreadSafeMap) {},
			threads:     1,
			expectedVal: nil,
			expectedOk:  false,
		},
		{
			name:        "delete not found concurrent",
			key:         "key",
			before:      func(m ThreadSafeMap) {},
			threads:     100,
			expectedVal: nil,
			expectedOk:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := NewThreadSafeMap()
			tt.before(m)

			wg := sync.WaitGroup{}
			wg.Add(tt.threads)
			for i := 0; i < tt.threads; i++ {
				go func() {
					defer wg.Done()
					m.Delete(tt.key)
				}()
			}

			wg.Wait()
			val, ok := m.(*ThreadSafeMapImpl).internalMap[tt.key]
			require.Equal(t, tt.expectedVal, val)
			require.Equal(t, tt.expectedOk, ok)
		})
	}
}

func TestThreadSafeMap_Keys(t *testing.T) {
	tests := []struct {
		name        string
		before      func(m ThreadSafeMap)
		threads     int
		expectedVal []string
	}{
		{
			name: "keys",
			before: func(m ThreadSafeMap) {
				m.(*ThreadSafeMapImpl).internalMap["key1"] = "value1"
				m.(*ThreadSafeMapImpl).internalMap["key2"] = "value2"
				m.(*ThreadSafeMapImpl).internalMap["key3"] = "value3"
			},
			threads:     1,
			expectedVal: []string{"key1", "key2", "key3"},
		},
		{
			name: "keys concurrent",
			before: func(m ThreadSafeMap) {
				m.(*ThreadSafeMapImpl).internalMap["key1"] = "value1"
				m.(*ThreadSafeMapImpl).internalMap["key2"] = "value2"
				m.(*ThreadSafeMapImpl).internalMap["key3"] = "value3"
			},
			threads:     100,
			expectedVal: []string{"key1", "key2", "key3"},
		},
		{
			name:        "keys empty",
			before:      func(m ThreadSafeMap) {},
			threads:     1,
			expectedVal: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := NewThreadSafeMap()
			tt.before(m)

			wg := sync.WaitGroup{}
			wg.Add(tt.threads)
			for i := 0; i < tt.threads; i++ {
				go func() {
					defer wg.Done()
					keys := m.Keys()
					require.ElementsMatch(t, tt.expectedVal, keys)
				}()
			}

			wg.Wait()
		})
	}
}

func TestThreadSafeMap_Values(t *testing.T) {
	tests := []struct {
		name        string
		before      func(m ThreadSafeMap)
		threads     int
		expectedVal []interface{}
	}{
		{
			name: "values",
			before: func(m ThreadSafeMap) {
				m.(*ThreadSafeMapImpl).internalMap["key1"] = "value1"
				m.(*ThreadSafeMapImpl).internalMap["key2"] = "value2"
				m.(*ThreadSafeMapImpl).internalMap["key3"] = "value3"
			},
			threads:     1,
			expectedVal: []interface{}{"value1", "value2", "value3"},
		},
		{
			name: "values concurrent",
			before: func(m ThreadSafeMap) {
				m.(*ThreadSafeMapImpl).internalMap["key1"] = "value1"
				m.(*ThreadSafeMapImpl).internalMap["key2"] = "value2"
				m.(*ThreadSafeMapImpl).internalMap["key3"] = "value3"
			},
			threads:     100,
			expectedVal: []interface{}{"value1", "value2", "value3"},
		},
		{
			name:        "values empty",
			before:      func(m ThreadSafeMap) {},
			threads:     1,
			expectedVal: []interface{}{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := NewThreadSafeMap()
			tt.before(m)

			wg := sync.WaitGroup{}
			wg.Add(tt.threads)
			for i := 0; i < tt.threads; i++ {
				go func() {
					defer wg.Done()
					values := m.Values()
					require.ElementsMatch(t, tt.expectedVal, values)
				}()
			}

			wg.Wait()
		})
	}
}

func TestThreadSafeMap_ForEach(t *testing.T) {
	tests := []struct {
		name        string
		before      func(m ThreadSafeMap)
		threads     int
		expectedVal map[string]interface{}
	}{
		{
			name: "for each",
			before: func(m ThreadSafeMap) {
				m.(*ThreadSafeMapImpl).internalMap["key1"] = "value1"
				m.(*ThreadSafeMapImpl).internalMap["key2"] = "value2"
				m.(*ThreadSafeMapImpl).internalMap["key3"] = "value3"
			},
			threads: 1,
			expectedVal: map[string]interface{}{
				"key1": "value1",
				"key2": "value2",
				"key3": "value3",
			},
		},
		{
			name: "for each concurrent",
			before: func(m ThreadSafeMap) {
				m.(*ThreadSafeMapImpl).internalMap["key1"] = "value1"
				m.(*ThreadSafeMapImpl).internalMap["key2"] = "value2"
				m.(*ThreadSafeMapImpl).internalMap["key3"] = "value3"
			},
			threads: 100,
			expectedVal: map[string]interface{}{
				"key1": "value1",
				"key2": "value2",
				"key3": "value3",
			},
		},
		{
			name:        "for each empty",
			before:      func(m ThreadSafeMap) {},
			threads:     1,
			expectedVal: map[string]interface{}{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := NewThreadSafeMap()
			tt.before(m)

			wg := sync.WaitGroup{}
			wg.Add(tt.threads)
			for i := 0; i < tt.threads; i++ {
				go func() {
					defer wg.Done()

					m.ForEach(func(key string, val interface{}) {
						require.Equal(t, tt.expectedVal[key], val)
					})
				}()
			}

			wg.Wait()
		})
	}
}
