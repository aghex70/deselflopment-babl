package entry

import (
	"context"
	"github.com/aghex70/deselflopment-babl/internal/core/domain"
	"github.com/aghex70/deselflopment-babl/internal/core/ports"
	entryStore "github.com/aghex70/deselflopment-babl/internal/stores/entry"
	"github.com/aghex70/deselflopment-babl/pkg"
	"time"
)

type Service struct {
	entryRepository *entryStore.GormRepository
}

func (s Service) Create(ctx context.Context, r ports.CreateEntryRequest) (domain.Entry, error) {
	uuid := pkg.GenerateUUID()
	ee := domain.Entry{
		Id:          uuid,
		Name:        r.Name,
		EventType:   domain.EventType(r.EventType),
		EventDate:   r.EventDate,
		Origin:      r.Origin,
		Description: r.Description,
		Duration:    time.Duration(r.Duration),
		Score:       r.Score,
		Positive:    r.Positive,
	}
	ne, err := s.entryRepository.Create(ctx, ee)
	if err != nil {
		return domain.Entry{}, err
	}

	return ne, nil
}

func (s Service) Update(ctx context.Context, r ports.UpdateEntryRequest) (domain.Entry, error) {
	ee := domain.Entry{
		Id:   r.Id,
		Name: r.Name,
	}
	ue, err := s.entryRepository.Update(ctx, ee)
	if err != nil {
		return domain.Entry{}, err
	}
	return ue, nil
}

func (s Service) Get(ctx context.Context, uuid string) (domain.Entry, error) {
	ee, err := s.entryRepository.GetById(ctx, uuid)
	if err != nil {
		return domain.Entry{}, err
	}
	return ee, nil
}

func (s Service) Delete(ctx context.Context, uuid string) error {
	err := s.entryRepository.Delete(ctx, uuid)
	if err != nil {
		return err
	}
	return nil
}

func (s Service) List(ctx context.Context) ([]domain.Entry, error) {
	es, err := s.entryRepository.List(ctx)
	if err != nil {
		return []domain.Entry{}, err
	}
	return es, nil
}

func NewService(er *entryStore.GormRepository) (Service, error) {
	return Service{
		entryRepository: er,
	}, nil
}
