package locku

import "sync"

type Cond interface {
	Lock()
	Unlock()
	Wait()
}

func NewCond() Cond {
	return &cond{
		cond: sync.NewCond(&sync.Mutex{}),
	}
}

type cond struct {
	cond *sync.Cond
}

func (c *cond) Lock() {
	c.cond.L.Lock()
}

func (c *cond) Unlock() {
	c.cond.L.Unlock()
	c.cond.Signal()
}

func (c *cond) Wait() {
	c.cond.Wait()
}
