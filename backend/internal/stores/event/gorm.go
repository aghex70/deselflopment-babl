package event

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

type Event struct {
	Id            string     `gorm:"primaryKey;column:id"`
	Name          string     `gorm:"column:name"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type Tabler interface {
	TableName() string
}

// TableName overrides the table name
func (Event) TableName() string {
	return "deselflopment-babl_events"
}

func (gr *GormRepository) Create(ctx context.Context, e domain.Event) (domain.Event, error) {
	ne := fromDto(e)
	result := gr.DB.Create(&ne)
	if result.Error != nil {
		return domain.Event{}, result.Error
	}
	return ne.ToDto(), nil
}

func (gr *GormRepository) Update(ctx context.Context, e domain.Event) (domain.Event, error) {
	ue := fromDto(e)
	result := gr.DB.Model(&ue).Where(Event{Id: ue.Id}).Updates(map[string]interface{}{
		"name": ue.Name,
	})

	if result.RowsAffected == 0 {
		return domain.Event{}, gorm.ErrRecordNotFound
	}

	if result.Error != nil {
		return domain.Event{}, result.Error
	}
	return ue.ToDto(), nil
}

func (gr *GormRepository) GetById(ctx context.Context, uuid string) (domain.Event, error) {
	var ee Event
	result := gr.DB.Where(&Event{Id: uuid}).First(&ee)
	if result.Error != nil {
		return domain.Event{}, result.Error
	}
	return ee.ToDto(), nil
}

func (gr *GormRepository) List(ctx context.Context) ([]domain.Event, error) {
	var es []Event
	var ees []domain.Event
	result := gr.DB.Find(&es)
	if result.Error != nil {
		return []domain.Event{}, result.Error
	}

	for _, e := range es {
		ee := e.ToDto()
		ees = append(ees, ee)
	}
	return ees, nil
}

func (gr *GormRepository) Delete(ctx context.Context, uuid string) error {
	ee := Event{Id: uuid}
	result := gr.DB.Delete(&ee)
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

func (ee Event) ToDto() domain.Event {
	return domain.Event{
		Name: 		  ee.Name,
		CreatedAt:    ee.CreatedAt,
		Id:           ee.Id,
		UpdatedAt:     ee.UpdatedAt,
	}
}

func fromDto(ee domain.Event) Event {
	return Event{
		Name: 		  ee.Name,
		CreatedAt:    ee.CreatedAt,
		Id:           ee.Id,
		UpdatedAt:    ee.UpdatedAt,
	}
}
