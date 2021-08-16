package syncutil

import "sync"

type Counter struct {
	Name string

	// エクスポートされないフィールドがある
	// 場合は、空行を入れることが多い
	// ミューテックスを利用する際は対象となる
	// フィールドらの先頭で定義することが多い
	m     sync.RWMutex
	count int
}

func (c *Counter) Increment() int {
	c.m.Lock()
	defer c.m.Unlock()
	c.count++
	return c.count
}

func (c *Counter) View() int {
	c.m.RLock()
	defer c.m.RUnlock()
	return c.count
}
