package usecase

import (
	"context"

	"github.com/cloudwego/kitex-layout/internal/entity"
)

type Repository interface {
	GetByName(ctx context.Context, name string) (*entity.Biology, error)
	Create(ctx context.Context, biology *entity.Biology) error
}
