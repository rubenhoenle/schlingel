package persistence

import (
	"github.com/google/uuid"
	"github.com/jgero/schlingel/api"
	"github.com/jgero/schlingel/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type OrmPersistence struct {
	db *gorm.DB
}

func NewOrmPersistence() api.Persistence {
	db, err := gorm.Open(sqlite.Open("schlingel.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&model.SchlingelFile{})

	return &OrmPersistence{db: db}
}

func (p *OrmPersistence) CreateFile(file model.SchlingelFile) error {
	result := p.db.Create(&file)
	return result.Error
}

func (p *OrmPersistence) GetFileByUuid(fileUuid uuid.UUID) (*model.SchlingelFile, error) {
	var file model.SchlingelFile
	p.db.First(&file, "uuid = ?", fileUuid.String())
	return &file, nil
}
