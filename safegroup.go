package safegroup

import (
	"context"
	"fmt"
	"runtime/debug"

	"golang.org/x/sync/errgroup"
)

type SafeGroup struct {
	g *errgroup.Group
}

func safeWrap(f func() error) func() error {
	return func() (retErr error) {
		defer func() {
			if expt := recover(); expt != nil {
				retErr = &PanicError{
					Expt:  expt,
					Stack: debug.Stack(),
				}
			}
		}()
		return f()
	}
}

func New() *SafeGroup {
	g := new(errgroup.Group)
	return &SafeGroup{g: g}
}

func WithContext(ctx context.Context) (*SafeGroup, context.Context) {
	eg, ctx := errgroup.WithContext(ctx)
	return &SafeGroup{g: eg}, ctx
}

func (sg *SafeGroup) Go(f func() error) {
	sg.g.Go(safeWrap(f))
}

func (sg *SafeGroup) Wait() error {
	return sg.g.Wait()
}

func (sg *SafeGroup) TryGo(f func() error) bool {
	return sg.g.TryGo(safeWrap(f))
}

func (sg *SafeGroup) SetLimit(n int) {
	sg.g.SetLimit(n)
}

// MARK: PanicError
type PanicError struct {
	Expt  any
	Stack []byte
}

var _ error = &PanicError{}

func (pe *PanicError) Error() string {
	return fmt.Sprintf("panic occurs: %v\n%s", pe.Expt, pe.Stack)
}
