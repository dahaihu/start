package main

import (
	"context"
	"fmt"
	"sync"
)

type Group struct {
	ch      chan func(context.Context) error
	chs     []func(context.Context) error
	wg      sync.WaitGroup
	cancel  func()
	ctx     context.Context
	errOnce sync.Once
	runOnce sync.Once
	err     error
}

func (g *Group) MaxGroups(n int) {
	if n <= 0 {
		panic(fmt.Errorf("invalid groups %d", n))
	}
	g.runOnce.Do(func() {
		g.ch = make(chan func(context.Context) error, n)
		for i := 0; i < n; i++ {
			go func() {
				for f := range g.ch {
					g.do(f)
				}
			}()
		}
	})
}

func (g *Group) do(f func(ctx context.Context) error) {
	ctx := g.ctx
	if ctx == nil {
		ctx = context.Background()
	}
	var err error
	defer func() {
		if err != nil {
			g.errOnce.Do(func() {
				g.err = err
				if g.cancel != nil {
					g.cancel()
				}
			})
		}
		g.wg.Done()
	}()
	err = f(ctx)
}

func (g *Group) Go(f func(ctx context.Context) error) {
	g.wg.Add(1)
	if g.ch != nil {
		select {
		case g.ch <- f:
		default:
			g.chs = append(g.chs, f)
		}
		return
	}
	go g.do(f)
}

func (g *Group) Wait() error {
	if g.ch != nil {
		for _, f := range g.chs {
			g.ch <- f
		}
	}
	g.wg.Wait()
	if g.ch != nil {
		close(g.ch)
	}
	if g.cancel != nil {
		g.cancel()
	}
	return g.err
}
