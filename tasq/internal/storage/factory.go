package storage

import "errors"

var ErrInvalidStorageEnrgine = errors.New("failed to specify a valid storage engine type")

type StorageEngineType string

const (
	File StorageEngineType = "csv"
	SQL  StorageEngineType = "sql"
)

// Available engine types
func SelectEngine(engineType StorageEngineType) (Storage, error) {
	switch engineType {
	case File:
		return newCsvEngine("/Users/leandro.landmeyer/Repos/Github/goprojects/tasq/tasks.csv"), nil
	default:
		return nil, ErrInvalidStorageEnrgine
	}
}
