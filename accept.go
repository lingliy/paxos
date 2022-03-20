package paxos

import (
	"context"
	"github/lingliy/paxos/basic"
	"sync"
)

type accept struct {
	basic.UnimplementedBasicServer
	state basic.Response
	mu    sync.Mutex
}

func InitAccept() *accept {
	return &accept{}
}

func (a *accept) Prepare(ctx context.Context, req *basic.Request) (*basic.Response, error) {
	a.mu.Lock()
	defer a.mu.Unlock()
	ret := a.state
	if req.Id > a.state.LastId {
		a.state.LastId = req.Id
	}
	return &ret, nil
}

func (a *accept) Accept(ctx context.Context, req *basic.Response2) (*basic.Request, error) {
	a.mu.Lock()
	defer a.mu.Unlock()
	if req.Id >= a.state.LastId {
		a.state.Id = req.Id
		a.state.State = req.State
		a.state.LastId = req.Id
	}
	return &basic.Request{Id: a.state.LastId}, nil
}
