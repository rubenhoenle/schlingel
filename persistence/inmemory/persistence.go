package inmemory

import (
	"github.com/google/uuid"
	"github.com/jgero/schlingel/api"
	"github.com/jgero/schlingel/model"
)

// implements the Persistence interface
type InMemoryPersistence struct {
	files []model.SchlingelFile
}

func NewInMemoryPersistence() *InMemoryPersistence {
	return &InMemoryPersistence{files: []model.SchlingelFile{}}
}

func (p *InMemoryPersistence) NewTx() (api.TransactionHandler, error) {
	return &InMemoryTransactionHandler{persistence: p}, nil
}

func (p *InMemoryPersistence) GetFileByUuid(fileUuid uuid.UUID) (*model.SchlingelFile, error) {
	for _, f := range p.files {
		if f.UUID == fileUuid {
			return &f, nil
		}
	}
	return nil, nil
}
