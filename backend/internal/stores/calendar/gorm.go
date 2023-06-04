package calendar

import (
	"context"
	"database/sql"
	"github.com/aghex70/deselflopment-babl/internal/core/domain"
	store "github.com/aghex70/deselflopment-babl/internal/stores/user"
	"gorm.io/gorm"
	"log"
	"time"
)

type GormRepository struct {
	*gorm.DB
	SqlDb  *sql.DB
	logger *log.Logger
}

type Calendar struct {
	Id        string `gorm:"primaryKey;column:id"`
	Name      string
	UserID    uint
	User      store.User `gorm:"foreignkey:UserID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Tabler interface {
	TableName() string
}

// TableName overrides the table name
func (Calendar) TableName() string {
	return "babl_calendars"
}

func (gr *GormRepository) Create(ctx context.Context, c domain.Calendar) (domain.Calendar, error) {
	nc := fromDto(c)
	result := gr.DB.Create(&nc)
	if result.Error != nil {
		return domain.Calendar{}, result.Error
	}
	return nc.ToDto(), nil
}

func (gr *GormRepository) Update(ctx context.Context, c domain.Calendar) (domain.Calendar, error) {
	uc := fromDto(c)
	result := gr.DB.Model(&uc).Where(Calendar{Id: uc.Id}).Updates(map[string]interface{}{
		"name": uc.Name,
	})

	if result.RowsAffected == 0 {
		return domain.Calendar{}, gorm.ErrRecordNotFound
	}

	if result.Error != nil {
		return domain.Calendar{}, result.Error
	}
	return uc.ToDto(), nil
}

func (gr *GormRepository) GetById(ctx context.Context, uuid string) (domain.Calendar, error) {
	var cc Calendar
	result := gr.DB.Where(&Calendar{Id: uuid}).First(&cc)
	if result.Error != nil {
		return domain.Calendar{}, result.Error
	}
	return cc.ToDto(), nil
}

func (gr *GormRepository) List(ctx context.Context) ([]domain.Calendar, error) {
	var cs []Calendar
	var ccs []domain.Calendar
	result := gr.DB.Find(&cs)
	if result.Error != nil {
		return []domain.Calendar{}, result.Error
	}

	for _, c := range cs {
		cc := c.ToDto()
		ccs = append(ccs, cc)
	}
	return ccs, nil
}

func (gr *GormRepository) Delete(ctx context.Context, uuid string) error {
	cc := Calendar{Id: uuid}
	result := gr.DB.Delete(&cc)
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

func (cc Calendar) ToDto() domain.Calendar {
	return domain.Calendar{
		Name:      cc.Name,
		CreatedAt: cc.CreatedAt,
		Id:        cc.Id,
		UpdatedAt: cc.UpdatedAt,
	}
}

func fromDto(cc domain.Calendar) Calendar {
	return Calendar{
		Name:      cc.Name,
		CreatedAt: cc.CreatedAt,
		Id:        cc.Id,
		UpdatedAt: cc.UpdatedAt,
	}
}
