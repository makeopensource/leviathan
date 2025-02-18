package models

import "gorm.io/gorm"

type LabModel struct {
	gorm.Model
	LabName        string
	ImageTag       string
	DockerFilePath string
}
