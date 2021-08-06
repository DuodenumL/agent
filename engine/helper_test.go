package engine

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBackoffRetry(t *testing.T) {
	i := 0
	f := func() error {
		i++
		if i < 3 {
			return errors.New("xxx")
		}
		return nil
	}
	assert.Nil(t, backoffRetry(context.Background(), 10, f))
}
