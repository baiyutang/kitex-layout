//go:build wireinject
// +build wireinject

package main

import (
	"github.com/cloudwego/kitex-layout/internal/config"
	"github.com/cloudwego/kitex-layout/internal/infrastructure/repository"
	"github.com/cloudwego/kitex-layout/internal/usecase"
	"github.com/cloudwego/kitex-layout/kitex_gen"
	"github.com/google/wire"
)

func initHandler(*config.DB) kitex_gen.World {
	panic(wire.Build(
		repository.SQLProviderSet,
		usecase.ProviderSet,
	))
}
