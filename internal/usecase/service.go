package usecase

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/cloudwego/kitex-layout/internal/entity"
	"github.com/cloudwego/kitex-layout/kitex_gen"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/google/wire"
)

// Service implements the last service interface defined in the IDL.
type Service struct {
	repo Repository
}

var ProviderSet = wire.NewSet(NewService)

func NewService(repo Repository) kitex_gen.World {
	return &Service{
		repo: repo,
	}
}

// Hello implements the World interface.
func (s *Service) Hello(ctx context.Context, req *kitex_gen.HelloReq) (resp *kitex_gen.HelloResp, err error) {
	biology, err := s.repo.GetByName(ctx, req.Name)
	if err != nil {
		return nil, err
	}
	klog.Infof("we found the biology named %s,whose ID is %d", biology.Name, biology.ID)
	return &kitex_gen.HelloResp{Message: fmt.Sprintf("hi!%s!!", req.Name)}, nil
}

func (s *Service) Evolve(ctx context.Context, req *kitex_gen.EvolveReq) (res *kitex_gen.EvolveResp, err error) {
	b := entity.NewBiology()
	b.Name = fmt.Sprintf("%d", rand.Int31())
	if err = s.repo.Create(ctx, b); err != nil {
		return nil, err
	}
	return &kitex_gen.EvolveResp{Name: b.Name}, nil
}
