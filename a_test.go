package paxos

import (
	"sync"
	"testing"
)

var count = 1000
var client = 3

func TestHello(t *testing.T) {
	arr := Init(3)
	var wg sync.WaitGroup
	wg.Add(client)
	for i := 0; i < client; i++ {
		v := arr[i]
		go func(p *proposer) {
			for i := 0; i < count; i++ {
				p.Write(func(state int64) int64 { return state + 1 })
			}
			wg.Done()
		}(v)
	}
	wg.Wait()
	state := arr[0].Read()
	num := int64(client * count)
	if state != num {
		t.Fatalf("num:%v state:%v\n", num, state)
	}
}
