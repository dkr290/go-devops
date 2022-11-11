package models

import (
	"time"

	"gorm.io/gorm"
)

type Note struct {
	ID        uint64 `gorm:"primaryKey"`
	Name      string `gorm:"size:255"`
	Content   string `gorm:"type:text"`
	CreatedAt time.Time
	UpdatedAt time.Time      `gorm:"Index"`
	DeletedAt gorm.DeletedAt `gorm:"Index"`
}

func NotesAll() *[]Note {

	var notes []Note
	Repo.DB.Where("deleted_at is NULL").Order("updated_at desc").Find(&notes)
	return &notes
}

func NoteCreate(name, content string) *Note {

	noteEntry := Note{Name: name, Content: content}
	Repo.DB.Create(&noteEntry)
	return &noteEntry
}
