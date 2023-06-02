package user

import (
	"context"
	"github.com/aghex70/deselflopment-babl/internal/core/domain"
	"github.com/aghex70/deselflopment-babl/internal/core/ports"
	userStore "github.com/aghex70/deselflopment-babl/internal/stores/user"
)

type Service struct {
	userRepository   *userStore.GormRepository
}

func (s Service) Create(ctx context.Context, r ports.CreateUserRequest) (domain.User, error) {
	uu := domain.User{
		Name: 		  r.Name,
	}
	ne, err := s.userRepository.Create(ctx, uu)
	if err != nil {
		return domain.User{}, err
	}

	return ne, nil
}

func (s Service) Update(ctx context.Context, r ports.UpdateUserRequest) (domain.User, error) {
	uu := domain.User{
		Id:                 r.Id,
		Name: 		  		r.Name,
	}
	uu, err := s.userRepository.Update(ctx, uu)
	if err != nil {
		return domain.User{}, err
	}
	return uu, nil
}

func (s Service) Get(ctx context.Context, uuid string) (domain.User, error) {
	uu, err := s.userRepository.GetById(ctx, uuid)
	if err != nil {
		return domain.User{}, err
	}
	return uu, nil
}

func (s Service) Delete(ctx context.Context, uuid string) error {
	err := s.userRepository.Delete(ctx, uuid)
	if err != nil {
		return err
	}
	return nil
}

func (s Service) List(ctx context.Context) ([]domain.User, error) {
	us, err := s.userRepository.List(ctx)
	if err != nil {
		return []domain.User{}, err
	}
	return us, nil
}

func NewService(ur *userStore.GormRepository) (Service, error) {
	return Service{
		userRepository:      ur,
	}, nil
}
