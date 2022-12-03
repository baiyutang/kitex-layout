package main

import (
	"context"

	"github.com/cloudwego/kitex-layout/kitex_gen"
	"github.com/cloudwego/kitex-layout/kitex_gen/world"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
)

func main() {
	c, err := world.NewClient("world", client.WithHostPorts("127.0.0.1:8080"))
	if err != nil {
		klog.Error(err)
	}
	rsp, err := c.Evolve(context.Background(), &kitex_gen.EvolveReq{})
	if err != nil {
		klog.Error(err)
	}
	reply, err := c.Hello(context.Background(), &kitex_gen.HelloReq{Name: rsp.Name})
	if err != nil {
		klog.Error(err)
	}
	klog.Infof("reply:%v", reply)
}
