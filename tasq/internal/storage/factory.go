package storage

import "errors"

// type StorageEngine struct {
// 	engine string
// }

type StorageEngineType string

const (
	File StorageEngineType = "csv"
)

// Available engine types
func SelectEngine(engineType StorageEngineType) (Storage, error) {
	switch engineType {
	case File:
		return newCsvEngine("/Users/leandro.landmeyer/Repos/Github/goprojects/tasq/tasks.csv"), nil
	default:
		return nil, errors.New("failed to specify a valid storage engine type")
	}
}
