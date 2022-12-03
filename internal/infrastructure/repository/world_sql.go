package repository

import (
	"context"
	"fmt"

	"github.com/cloudwego/kitex-layout/internal/config"
	"github.com/cloudwego/kitex-layout/internal/entity"
	"github.com/cloudwego/kitex-layout/internal/usecase"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type WorldRepository struct {
	db *gorm.DB
}

var SQLProviderSet = wire.NewSet(NewGorm, NewWorldSQL)

var _ usecase.Repository = (*WorldRepository)(nil)

func NewGorm(config *config.DB) *gorm.DB {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Username,
		config.Password,
		config.Host,
		config.Database,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	if err = db.AutoMigrate(&entity.Biology{}); err != nil {
		panic("failed to connect database")
	}
	return db
}

func NewWorldSQL(db *gorm.DB) usecase.Repository {
	return &WorldRepository{
		db: db,
	}
}

func (repo *WorldRepository) GetByName(_ context.Context, name string) (*entity.Biology, error) {
	obj := &entity.Biology{}
	result := repo.db.Where(&entity.Biology{Name: name}).First(obj)
	if result.Error != nil {
		return nil, result.Error
	}
	return &entity.Biology{
		ID:       obj.ID,
		Kingdom:  obj.Kingdom,
		Division: obj.Division,
		Class:    obj.Class,
		Order:    obj.Order,
		Family:   obj.Family,
		Genus:    obj.Genus,
		Species:  obj.Species,
		Name:     obj.Name,
	}, nil
}

func (repo *WorldRepository) Create(_ context.Context, biology *entity.Biology) error {
	result := repo.db.Create(biology)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
