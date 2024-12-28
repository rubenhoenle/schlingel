package inmemory

import "github.com/jgero/schlingel/model"

type InMemoryTransactionHandler struct {
	persistence *InMemoryPersistence
}

func (t *InMemoryTransactionHandler) Commit() error {
	// nothing to do here
	return nil
}

func (t *InMemoryTransactionHandler) Rollback() error {
	// nothing to do here
	return nil
}

func (t *InMemoryTransactionHandler) CreateFile(file model.SchlingelFile) error {
	t.persistence.files = append(t.persistence.files, file)
	return nil
}
