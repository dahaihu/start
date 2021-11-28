package cache

import "sync"

type caller struct {
	val interface{}
	err error
	wg  sync.WaitGroup
}

type Group struct {
	g  map[string]*caller
	mu sync.Mutex
}

func (g *Group) Do(key string, fn func() (interface{}, error)) (interface{}, error) {
	g.mu.Lock()
	if g.g == nil {
		g.g = make(map[string]*caller)
	}
	if c, ok := g.g[key]; ok {
		g.mu.Unlock()
		c.wg.Wait()
		return c.val, c.err
	}
	c := new(caller)
	c.wg.Add(1)
	g.g[key] = c
	g.mu.Unlock()

	c.val, c.err = fn()
	c.wg.Done()
	g.mu.Lock()
	delete(g.g, key)
	g.mu.Unlock()
	return c.val, c.err
}
