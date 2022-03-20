package paxos

import (
	"context"
	"fmt"
	"github/lingliy/paxos/basic"
	"time"

	"github.com/godruoyi/go-snowflake"
	"google.golang.org/grpc"
)

type proposer struct {
	arr []basic.BasicClient
}

func InitProposer() *proposer {
	return &proposer{}
}

func (p *proposer) AddNode(conn *grpc.ClientConn) {
	c := basic.NewBasicClient(conn)
	p.arr = append(p.arr, c)
}

func (p *proposer) phase1(id uint64) (state int64, isChosen bool, idInvalid bool) {
	quorum := len(p.arr)/2 + 1
	choseNum := quorum
	last := &basic.Response{}
	for _, c := range p.arr {
		ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*10)
		//ctx := context.Background()
		resp, err := c.Prepare(ctx, &basic.Request{Id: id})
		cancel()
		if err != nil {
			continue
		}
		if resp.LastId > id {
			idInvalid = true
			return
		}
		quorum--
		if resp.LastId == resp.Id {
			choseNum--
		}
		if resp.Id > last.Id {
			last = resp
		}
	}
	if choseNum > 0 {
		isChosen = true
	}
	if quorum <= 0 {
		state = last.State
		return
	}
	idInvalid = true
	return
}

func (p *proposer) phase2(resp *basic.Response2) bool {
	quorum := len(p.arr)/2 + 1
	for _, c := range p.arr {
		ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*10)
		//ctx := context.Background()
		id, err := c.Accept(ctx, resp)
		cancel()
		if err != nil {
			continue
		}
		if id == nil {
			continue
		}
		if id.Id == resp.Id {
			quorum--
		}

	}
	if quorum != len(p.arr)/2+1 {
		return true
	}
	return false
}

func (p *proposer) Write(cmd func(state int64) int64) {
	id := snowflake.ID()
	state, isChosen, invalid := p.phase1(id)
	if invalid {
		p.Write(cmd)
		return
	}
	if !isChosen {
		state = cmd(state)
	}
	fmt.Println("state ->", state, " id-> ", id)
	if p.phase2(&basic.Response2{Id: id, State: state}) && !isChosen {
		return
	}
	p.Write(cmd)
}

func (p *proposer) Read() int64 {
	quorum := len(p.arr)/2 + 1
	lastId := uint64(0)
	state := int64(0)
	for {
		for _, c := range p.arr {
			ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*10)
			//ctx := context.Background()
			resp, err := c.Prepare(ctx, &basic.Request{Id: 0})
			cancel()
			if err != nil {
				continue
			}
			if resp.Id > lastId {
				lastId = resp.Id
				state = resp.State
			}
			quorum--
			if quorum == 0 {
				return state
			}
		}
	}
}
