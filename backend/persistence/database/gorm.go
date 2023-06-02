package database

import (
	"fmt"
	"github.com/aghex70/deselflopment-babl/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewGormDB(cfg config.DatabaseConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s?parseTime=true", cfg.User, cfg.Password, cfg.Net, cfg.Host, cfg.Port, cfg.Name)
	//dsn = "deselflopment-babluser:deselflopment-bablpw@tcp(db:11306)/deselflopment-babl"
	gormDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return gormDB, nil
}
