package models

import "gorm.io/gorm"

type LabModel struct {
	gorm.Model
	LabName        string
	GraderFilename string
	GraderFile     []byte
	MakeFilename   string
	MakeFile       []byte
}
