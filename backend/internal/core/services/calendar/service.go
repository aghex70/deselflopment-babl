package calendar

import (
	"context"
	"github.com/aghex70/deselflopment-babl/internal/core/domain"
	"github.com/aghex70/deselflopment-babl/internal/core/ports"
	calendarStore "github.com/aghex70/deselflopment-babl/internal/stores/calendar"
)

type Service struct {
	calendarRepository   *calendarStore.GormRepository
}

func (s Service) Create(ctx context.Context, r ports.CreateCalendarRequest) (domain.Calendar, error) {
	cc := domain.Calendar{
		Name: 		  r.Name,
	}
	ne, err := s.calendarRepository.Create(ctx, cc)
	if err != nil {
		return domain.Calendar{}, err
	}

	return ne, nil
}

func (s Service) Update(ctx context.Context, r ports.UpdateCalendarRequest) (domain.Calendar, error) {
	cc := domain.Calendar{
		Id:                 r.Id,
		Name: 		  		r.Name,
	}
	uc, err := s.calendarRepository.Update(ctx, cc)
	if err != nil {
		return domain.Calendar{}, err
	}
	return uc, nil
}

func (s Service) Get(ctx context.Context, uuid string) (domain.Calendar, error) {
	cc, err := s.calendarRepository.GetById(ctx, uuid)
	if err != nil {
		return domain.Calendar{}, err
	}
	return cc, nil
}

func (s Service) Delete(ctx context.Context, uuid string) error {
	err := s.calendarRepository.Delete(ctx, uuid)
	if err != nil {
		return err
	}
	return nil
}

func (s Service) List(ctx context.Context) ([]domain.Calendar, error) {
	cs, err := s.calendarRepository.List(ctx)
	if err != nil {
		return []domain.Calendar{}, err
	}
	return cs, nil
}

func NewService(cr *calendarStore.GormRepository) (Service, error) {
	return Service{
		calendarRepository:      cr,
	}, nil
}
