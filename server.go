package paxos

import (
	"fmt"
	"github/lingliy/paxos/basic"
	"net"
	"time"

	"google.golang.org/grpc"
)

var GRPCBasePort = 9000

func Init(num int) []*proposer {
	for i := 0; i < num; i++ {
		go initServer(fmt.Sprintf(":%d", GRPCBasePort+i))
	}
	time.Sleep(time.Second * 2)
	var address []string
	for i := 0; i < num; i++ {
		address = append(address, fmt.Sprintf("localhost:%d", i+GRPCBasePort))
	}
	var ret []*proposer
	for i := 0; i < num; i++ {
		ret = append(ret, initClient(uint16(i), address))
	}
	return ret
}

func initServer(address string) {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	basic.RegisterBasicServer(s, InitAccept())
	err = s.Serve(lis)
	if err != nil {
		panic(err)
	}
}

func initClient(machineID uint16, address []string) *proposer {
	prop := InitProposer()
	for _, v := range address {
		conn, err := grpc.Dial(v, grpc.WithInsecure())
		if err != nil {
			panic(err)
		}
		prop.AddNode(conn)
	}
	return prop
}
