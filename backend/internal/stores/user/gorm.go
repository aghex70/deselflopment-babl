package user

import (
	"context"
	"database/sql"
	"github.com/aghex70/deselflopment-babl/internal/core/domain"
	"gorm.io/gorm"
	"log"
	"time"
)

type GormRepository struct {
	*gorm.DB
	SqlDb  *sql.DB
	logger *log.Logger
}

type User struct {
	Id        string `gorm:"primaryKey;column:id"`
	Name      string
	Active    bool
	Admin     bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Tabler interface {
	TableName() string
}

// TableName overrides the table name
func (User) TableName() string {
	return "babl_users"
}

func (gr *GormRepository) Create(ctx context.Context, u domain.User) (domain.User, error) {
	nu := fromDto(u)
	result := gr.DB.Create(&nu)
	if result.Error != nil {
		return domain.User{}, result.Error
	}
	return nu.ToDto(), nil
}

func (gr *GormRepository) Update(ctx context.Context, u domain.User) (domain.User, error) {
	uu := fromDto(u)
	result := gr.DB.Model(&uu).Where(User{Id: uu.Id}).Updates(map[string]interface{}{
		"name": uu.Name,
	})

	if result.RowsAffected == 0 {
		return domain.User{}, gorm.ErrRecordNotFound
	}

	if result.Error != nil {
		return domain.User{}, result.Error
	}
	return uu.ToDto(), nil
}

func (gr *GormRepository) GetById(ctx context.Context, uuid string) (domain.User, error) {
	var uu User
	result := gr.DB.Where(&User{Id: uuid}).First(&uu)
	if result.Error != nil {
		return domain.User{}, result.Error
	}
	return uu.ToDto(), nil
}

func (gr *GormRepository) List(ctx context.Context) ([]domain.User, error) {
	var us []User
	var uus []domain.User
	result := gr.DB.Find(&us)
	if result.Error != nil {
		return []domain.User{}, result.Error
	}

	for _, u := range us {
		uu := u.ToDto()
		uus = append(uus, uu)
	}
	return uus, nil
}

func (gr *GormRepository) Delete(ctx context.Context, uuid string) error {
	uu := User{Id: uuid}
	result := gr.DB.Delete(&uu)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func NewGormRepository(db *gorm.DB) (*GormRepository, error) {
	return &GormRepository{
		DB: db,
	}, nil
}

func (uu User) ToDto() domain.User {
	return domain.User{
		Name:      uu.Name,
		CreatedAt: uu.CreatedAt,
		Id:        uu.Id,
		UpdatedAt: uu.UpdatedAt,
	}
}

func fromDto(uu domain.User) User {
	return User{
		Name:      uu.Name,
		CreatedAt: uu.CreatedAt,
		Id:        uu.Id,
		UpdatedAt: uu.UpdatedAt,
	}
}
