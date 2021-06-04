package database

import "context"

type IDatabase interface {
	Create(ctx context.Context, database, table string, entity interface{}) (string, error)
	GetByID(ctx context.Context, database, table string, entityId string, entity interface{}) error
	GetAll(ctx context.Context, database, table string, entities interface{}, limit int64, index int64) error
}

type Database struct {
}

func Build() IDatabase {
	return &Database{}
}

func (db *Database) Create(ctx context.Context, database, table string, entity interface{}) (string, error) {
	return "", nil
}

func (db *Database) GetByID(ctx context.Context, database, table string, entityId string, entity interface{}) error {
	return nil
}

func (db *Database) GetAll(ctx context.Context, database, table string, entities interface{}, limit int64, index int64) error {
	return nil
}
