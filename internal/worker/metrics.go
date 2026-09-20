package worker

import "sync/atomic"

type Counter struct {
	value atomic.Int64
}

func (c *Counter) Inc()         { c.value.Add(1) }
func (c *Counter) Value() int64 { return c.value.Load() }
func (c *Counter) Reset()       { c.value.Store(0) }
