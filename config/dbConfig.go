package config

import (
	"errors"
)

type DbConfig struct {
	MongoURI string
}

func (dbConfig *DbConfig) validate() error {
	if dbConfig.MongoURI == "" {
		return errors.New("MongoURI not supplied")
	}

	return nil
}
