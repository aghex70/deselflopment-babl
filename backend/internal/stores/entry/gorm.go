package entry

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

type Entry struct {
	Id          string `gorm:"primaryKey;column:id"`
	Name        string
	EventType   string
	EventDate   time.Time
	Origin      string
	Description string
	Duration    time.Duration
	Score       float32
	Positive    *bool
	ToImprove   *bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Tabler interface {
	TableName() string
}

// TableName overrides the table name
func (Entry) TableName() string {
	return "babl_entries"
}

func (gr *GormRepository) Create(ctx context.Context, e domain.Entry) (domain.Entry, error) {
	ne := fromDto(e)
	result := gr.DB.Create(&ne)
	if result.Error != nil {
		return domain.Entry{}, result.Error
	}
	return ne.ToDto(), nil
}

func (gr *GormRepository) Update(ctx context.Context, e domain.Entry) (domain.Entry, error) {
	ue := fromDto(e)
	result := gr.DB.Model(&ue).Where(Entry{Id: ue.Id}).Updates(map[string]interface{}{
		"name": ue.Name,
	})

	if result.RowsAffected == 0 {
		return domain.Entry{}, gorm.ErrRecordNotFound
	}

	if result.Error != nil {
		return domain.Entry{}, result.Error
	}
	return ue.ToDto(), nil
}

func (gr *GormRepository) GetById(ctx context.Context, uuid string) (domain.Entry, error) {
	var ee Entry
	result := gr.DB.Where(&Entry{Id: uuid}).First(&ee)
	if result.Error != nil {
		return domain.Entry{}, result.Error
	}
	return ee.ToDto(), nil
}

func (gr *GormRepository) List(ctx context.Context) ([]domain.Entry, error) {
	var es []Entry
	var ees []domain.Entry
	result := gr.DB.Find(&es)
	if result.Error != nil {
		return []domain.Entry{}, result.Error
	}

	for _, e := range es {
		ee := e.ToDto()
		ees = append(ees, ee)
	}
	return ees, nil
}

func (gr *GormRepository) Delete(ctx context.Context, uuid string) error {
	ee := Entry{Id: uuid}
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

func (ee Entry) ToDto() domain.Entry {
	return domain.Entry{
		Name:      ee.Name,
		CreatedAt: ee.CreatedAt,
		Id:        ee.Id,
		UpdatedAt: ee.UpdatedAt,
	}
}

func fromDto(ee domain.Entry) Entry {
	return Entry{
		Name:      ee.Name,
		CreatedAt: ee.CreatedAt,
		Id:        ee.Id,
		UpdatedAt: ee.UpdatedAt,
	}
}
