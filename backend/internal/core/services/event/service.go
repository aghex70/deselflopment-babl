package event

import (
	"context"
	"github.com/aghex70/deselflopment-babl/internal/core/domain"
	"github.com/aghex70/deselflopment-babl/internal/core/ports"
	eventStore "github.com/aghex70/deselflopment-babl/internal/stores/event"
)

type Service struct {
	eventRepository   *eventStore.GormRepository
}

func (s Service) Create(ctx context.Context, r ports.CreateEventRequest) (domain.Event, error) {
	ee := domain.Event{
		Name: 		  r.Name,
	}
	ne, err := s.eventRepository.Create(ctx, ee)
	if err != nil {
		return domain.Event{}, err
	}

	return ne, nil
}

func (s Service) Update(ctx context.Context, r ports.UpdateEventRequest) (domain.Event, error) {
	ee := domain.Event{
		Id:                 r.Id,
		Name: 		  		r.Name,
	}
	ue, err := s.eventRepository.Update(ctx, ee)
	if err != nil {
		return domain.Event{}, err
	}
	return ue, nil
}

func (s Service) Get(ctx context.Context, uuid string) (domain.Event, error) {
	ee, err := s.eventRepository.GetById(ctx, uuid)
	if err != nil {
		return domain.Event{}, err
	}
	return ee, nil
}

func (s Service) Delete(ctx context.Context, uuid string) error {
	err := s.eventRepository.Delete(ctx, uuid)
	if err != nil {
		return err
	}
	return nil
}

func (s Service) List(ctx context.Context) ([]domain.Event, error) {
	es, err := s.eventRepository.List(ctx)
	if err != nil {
		return []domain.Event{}, err
	}
	return es, nil
}

func NewService(er *eventStore.GormRepository) (Service, error) {
	return Service{
		eventRepository:      er,
	}, nil
}
