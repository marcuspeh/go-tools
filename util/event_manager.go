package util

import (
	"fmt"
)

type EventManager[T interface{}] interface {
	Subscribe(id string, callback func(T)) error
	Unsubscribe(id string)
	NotifyCallbacks(id string, data T)
}

func NewEventManager[T interface{}]() EventManager[T] {
	return &eventManagerImpl[T]{
		callbacks: make(map[string]func(T)),
	}
}

type eventManagerImpl[T interface{}] struct {
	callbacks map[string]func(T)
}

func (c *eventManagerImpl[T]) Subscribe(id string, callback func(T)) error {
	if _, ok := c.callbacks[id]; ok {
		return fmt.Errorf("callback id already exists")
	}

	c.callbacks[id] = callback
	return nil
}

func (c *eventManagerImpl[T]) Unsubscribe(id string) {
	delete(c.callbacks, id)
}

func (c *eventManagerImpl[T]) NotifyCallbacks(id string, data T) {
	for _, callback := range c.callbacks {
		callback(data)
	}
}
