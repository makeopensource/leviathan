package models

import (
	"gorm.io/gorm"
	"time"
)

type Lab struct {
	gorm.Model
	Name            string `gorm:"unique"`
	JobTimeout      time.Duration
	ImageTag        string
	DockerFilePath  string
	JobLimits       MachineLimits `gorm:"embedded;embeddedPrefix:machine_limit_"`
	JobEntryCmd     string
	JobFilesDirPath string
}

// VerifyJobLimits checks if job limits are provided,
// and sets fields that are missing with default values
func (l *Lab) VerifyJobLimits() {
	if l.JobLimits.PidsLimit == 0 {
		l.JobLimits.PidsLimit = 100 // Default value
	}
	if l.JobLimits.NanoCPU == 0 {
		l.JobLimits.NanoCPU = 1 // Default value
	}
	if l.JobLimits.Memory == 0 {
		l.JobLimits.Memory = 512 // Default value in MB
	}
}

type MachineLimits struct {
	PidsLimit int64
	// NanoCPU will be multiplied by CPUQuota
	NanoCPU int64
	// Memory in MB
	Memory int64
}
