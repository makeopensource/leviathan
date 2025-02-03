package utils

import (
	"fmt"
	"github.com/hashicorp/golang-lru/v2"
	"github.com/makeopensource/leviathan/models"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

// cache to store grader files

type LabFilesCache struct {
	cache *lru.Cache[string, *models.LabModel]
	db    *gorm.DB
}

func NewLabFilesCache(db *gorm.DB) *LabFilesCache {
	cacheSize := viper.GetInt("lab_files_cache")
	if cacheSize <= 0 {
		cacheSize = 35 // default cache size
	}

	cache, err := lru.New[string, *models.LabModel](cacheSize)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to initialize Lab Files Cache")
	}

	return &LabFilesCache{
		cache: cache,
		db:    db,
	}
}

func (lf *LabFilesCache) GetFromCache(labname string) (*models.LabModel, error) {
	if value, found := lf.cache.Get(labname); found {
		return value, nil
	}

	// cache miss, get data from database and load it up
	var labModel models.LabModel
	res := lf.db.Where("labname = ?", labname).First(&labModel)
	if res.Error != nil {
		log.Error().Err(res.Error).Msg("Failed to get Lab Model from DB")
		return nil, fmt.Errorf("failed to get lab info from DB")
	}

	lf.AddToCache(labname, &labModel)
	return &labModel, nil
}

func (lf *LabFilesCache) AddToCache(labname string, entry *models.LabModel) {
	lf.cache.Add(labname, entry)
}

func (lf *LabFilesCache) InvalidateKey(labname string) {
	lf.cache.Remove(labname)
}
