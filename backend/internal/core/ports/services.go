package ports

import (
	"context"
	"github.com/aghex70/deselflopment-babl/internal/core/domain"
)

type EntryServicer interface {
	Create(ctx context.Context, r CreateEntryRequest) (domain.Entry, error)
	Update(ctx context.Context, r UpdateEntryRequest) (domain.Entry, error)
	Get(ctx context.Context, uuid string) (domain.Entry, error)
	Delete(ctx context.Context, uuid string) error
	List(ctx context.Context) ([]domain.Entry, error)
}

type UserServicer interface {
	Get(ctx context.Context, uuid string) (domain.User, error)
	Delete(ctx context.Context, uuid string) error
	List(ctx context.Context) ([]domain.User, error)
}
