package utils

import (
	"context"
	log "github.com/sirupsen/logrus"
	"time"
)

func BackoffRetry(ctx context.Context, maxInterval int64, f func() error) error {
	t := time.NewTimer(0)
	var err error
	// make sure to execute at least once
	maxInterval = Max(maxInterval, 2)
	for i := int64(1); i < maxInterval; i *= 2 {
		select {
		case <-t.C:
			if err = f(); err == nil {
				return nil
			}
			log.Debugf("[backoffRetry] will retry after %d seconds", i)
			t.Reset(time.Duration(i) * time.Second)
		case <-ctx.Done():
			return err
		}
	}
	return err
}
