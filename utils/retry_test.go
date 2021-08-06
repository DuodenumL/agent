package utils

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
		if i < 1 {
			return errors.New("xxx")
		}
		return nil
	}
	assert.Nil(t, BackoffRetry(context.Background(), 1, f))
	assert.EqualValues(t, 1, i)
}
