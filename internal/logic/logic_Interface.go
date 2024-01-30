package logic

import (
	"context"
	"github.com/Michael-Levitin/YellowPages/internal/dto"
)

type PagesLogicI interface {
	GetInfo(ctx context.Context, info dto.Info) (dto.Info, error)
	SetInfo(ctx context.Context, info dto.Info) (dto.Info, error)
}