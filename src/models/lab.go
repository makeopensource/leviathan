package models

import "gorm.io/gorm"

type LabModel struct {
	gorm.Model
	LabName        string
	GraderFilename string
	GraderFile     []byte `gorm:"-"` // This field will be ignored by GORM
	MakeFilename   string
	MakeFile       []byte `gorm:"-"`
}
