package safegroup_test

import (
	"errors"
	"testing"

	"github.com/akishichinibu/safegroup"
	"github.com/stretchr/testify/require"
)

func TestPanicRecover(t *testing.T) {
	sg := safegroup.New()

	sg.Go(func() error {
		panic("something went wrong")
	})

	err := sg.Wait()
	require.Error(t, err)

	var pErr *safegroup.PanicError
	require.True(t, errors.As(err, &pErr))
	t.Logf("Recovered panic: %v", pErr.Expt)
	t.Logf("Stack trace:\n%s", string(pErr.Stack))
}
