package ports

import (
	"context"
	"github.com/aghex70/deselflopment-babl/internal/core/domain"
)

type CalendarRepository interface {
	Create(ctx context.Context, a domain.Calendar) (domain.Calendar, error)
	Update(ctx context.Context, a domain.Calendar) (domain.Calendar, error)
	GetById(ctx context.Context, uuid string) (domain.Calendar, error)
	List(ctx context.Context) ([]domain.Calendar, error)
	Delete(ctx context.Context, uuid string) error
}

type EventRepository interface {
	Create(ctx context.Context, a domain.Event) (domain.Event, error)
	Update(ctx context.Context, a domain.Event) (domain.Event, error)
	GetById(ctx context.Context, uuid string) (domain.Event, error)
	List(ctx context.Context) ([]domain.Event, error)
	Delete(ctx context.Context, uuid string) error
}

type UserRepository interface {
	Create(ctx context.Context, a domain.User) (domain.User, error)
	Update(ctx context.Context, a domain.User) (domain.User, error)
	GetById(ctx context.Context, uuid string) (domain.User, error)
	List(ctx context.Context) ([]domain.User, error)
	Delete(ctx context.Context, uuid string) error
}

