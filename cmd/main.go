package main

import (
	"net"

	"github.com/cloudwego/kitex-layout/internal/config"
	"github.com/cloudwego/kitex-layout/kitex_gen/world"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/spf13/viper"
)

func main() {
	c := config.New()

	viper.SetConfigName("app")
	viper.AddConfigPath("../configs/")
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := viper.Unmarshal(&c); err != nil {
		panic(err)
	}
	c.Validate()

	hdlr := initHandler(c.DB)
	svr := world.NewServer(
		hdlr,
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: c.Server.Name}),
		server.WithServiceAddr(&net.TCPAddr{Port: c.Server.Port}),
	)

	if err := svr.Run(); err != nil {
		klog.Fatal(err.Error())
	}
}
