package util

import "sync"

type ThreadSafeMap interface {
	Get(key string) (interface{}, bool)
	Set(key string, value interface{})
	Update(key string, value interface{}, fn func(interface{}, interface{}) interface{})
	AddListItem(key string, value interface{})
	Delete(key string)
	Len() int
	Keys() []string
	Values() []interface{}
	ForEach(fn func(key string, value interface{}))
}

type ThreadSafeMapImpl struct {
	sync        sync.RWMutex
	internalMap map[string]interface{}
}

func NewThreadSafeMap() ThreadSafeMap {
	return &ThreadSafeMapImpl{
		internalMap: make(map[string]interface{}),
	}
}

func (m *ThreadSafeMapImpl) Get(key string) (interface{}, bool) {
	m.sync.Lock()
	defer m.sync.Unlock()

	val, ok := m.internalMap[key]
	return val, ok
}
func (m *ThreadSafeMapImpl) Set(key string, value interface{}) {
	m.sync.Lock()
	defer m.sync.Unlock()

	m.internalMap[key] = value
}

func (m *ThreadSafeMapImpl) Update(key string, value interface{}, fn func(interface{}, interface{}) interface{}) {
	m.sync.Lock()
	defer m.sync.Unlock()

	m.internalMap[key] = fn(m.internalMap[key], value)
}

func (m *ThreadSafeMapImpl) AddListItem(key string, value interface{}) {
	m.sync.Lock()
	defer m.sync.Unlock()

	if internalList, ok := m.internalMap[key]; ok {
		m.internalMap[key] = append(internalList.([]interface{}), value)
		return
	}
	m.internalMap[key] = []interface{}{value}
}

func (m *ThreadSafeMapImpl) Delete(key string) {
	m.sync.Lock()
	defer m.sync.Unlock()

	delete(m.internalMap, key)
}

func (m *ThreadSafeMapImpl) Len() int {
	m.sync.Lock()
	defer m.sync.Unlock()

	return len(m.internalMap)
}

func (m *ThreadSafeMapImpl) Keys() []string {
	m.sync.Lock()
	defer m.sync.Unlock()

	keys := make([]string, 0, len(m.internalMap))
	for k := range m.internalMap {
		keys = append(keys, k)
	}
	return keys
}

func (m *ThreadSafeMapImpl) Values() []interface{} {
	m.sync.Lock()
	defer m.sync.Unlock()

	values := make([]interface{}, 0, len(m.internalMap))
	for _, v := range m.internalMap {
		values = append(values, v)
	}
	return values
}

func (m *ThreadSafeMapImpl) ForEach(fn func(key string, value interface{})) {
	m.sync.Lock()
	defer m.sync.Unlock()

	for k, v := range m.internalMap {
		fn(k, v)
	}
}
