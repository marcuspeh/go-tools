package tools

import (
	"context"
	"fmt"

	"golang.org/x/sync/errgroup"
)

type ErrGroup interface {
	Run(ctx context.Context, fn func() error)
	Wait() error
}

type ErrGroupImpl struct {
	grp errgroup.Group
}

func NewErrGroup() ErrGroup {
	return &ErrGroupImpl{
		grp: errgroup.Group{},
	}
}

func (m *ErrGroupImpl) Run(ctx context.Context, fn func() error) {
	m.grp.Go(func() (err error) {
		defer func() {
			if err := recover(); err != nil {
				err = fmt.Errorf("panic occured %v", err)
			}
		}()

		return fn()
	})
}

func (m *ErrGroupImpl) Wait() error {
	return m.grp.Wait()
}
