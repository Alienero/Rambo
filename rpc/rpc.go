package rpc

import (
	"net"
	"net/rpc"
)

type Server interface {
	ListenAndServe(addr string) error
	Register(rcvr interface{}) error
	RegisterName(name string, rcvr interface{}) error
}

type CallResult struct {
	err error
}

type GobServer struct {
	server rpc.Server
}

func (s *GobServer) ListenAndServe(addr string) error {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	s.server.Accept(l)
	return nil
}

func (s *GobServer) Register(rcvr interface{}) error {
	return s.server.Register(rcvr)
}

func (s *GobServer) RegisterName(name string, rcvr interface{}) error {
	return s.server.RegisterName(name, rcvr)
}

type Client interface {
	Call(serviceMethod string, args interface{}, reply interface{}) error
	Close() error
}

type GobClient struct {
	c *rpc.Client
}

func NewGobClient(addr string) (Client, error) {
	c, err := rpc.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}
	return &GobClient{
		c: c,
	}, nil
}

func (c *GobClient) Close() error {
	return c.c.Close()
}

func (c *GobClient) Call(serviceMethod string, args interface{}, reply interface{}) error {
	return c.c.Call(serviceMethod, args, reply)
}
