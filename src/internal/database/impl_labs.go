package database

import (
	"github.com/makeopensource/leviathan/internal/labs"
	"gorm.io/gorm"
)

type LabDatabase struct {
	db *gorm.DB
}

func (l *LabDatabase) CreateLab(lab *labs.Lab) error {
	res := l.db.Save(lab)
	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (l *LabDatabase) DeleteLab(id uint) error {
	if res := l.db.Delete(&labs.Lab{}, id); res.Error != nil {
		return res.Error
	}

	return nil
}

func (l *LabDatabase) GetLab(id uint) (*labs.Lab, error) {
	var lab labs.Lab
	if res := l.db.Where("ID = ?", id).First(&lab); res.Error != nil {
		return nil, res.Error
	}

	return &lab, nil
}
