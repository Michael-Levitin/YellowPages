package database

import (
	"context"
	"github.com/Michael-Levitin/YellowPages/internal/dto"
)

type PagesDbI interface {
	GetInfo(ctx context.Context, info *dto.Info, page *dto.Page) (*[]dto.Info, error)
	SetInfo(ctx context.Context, info *dto.Info) (*dto.Info, error)
	DeleteInfo(ctx context.Context, info *dto.Info, page *dto.Page) (*[]dto.Info, error)
	UpdateInfo(ctx context.Context, info *dto.Info) (*dto.Info, error)
}
