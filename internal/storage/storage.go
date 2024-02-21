package storage

import (
	"gorm.io/gorm"
	"one-day-job/pkg/logger"
)

type Storage struct {
	db  *gorm.DB
	log *logger.Logger
}

func NewStorage(db *gorm.DB, log *logger.Logger) *Storage {
	return &Storage{
		db:  db,
		log: log,
	}
}
