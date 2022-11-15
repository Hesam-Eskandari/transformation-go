package repository

import (
	"encoding/json"
	"fmt"
	"github.com/Hesam-Eskandari/transformer/pkg/internal"
)

type DatabaseModels interface {
	GetRawRelationalModels() (string, error)
	GetRawCollections() (string, error)
	GetRawKeyValueModels() (interface{}, error)
}

type databaseModels struct {
	rawJSON     string
	modelsMap   map[string]interface{}
	tables      string
	collections string
	keyValues   string
}

func NewDatabaseModels(rawJSON string) DatabaseModels {
	return &databaseModels{
		rawJSON:     rawJSON,
		tables:      internal.Tables,
		collections: internal.Collections,
		keyValues:   internal.KeyValues,
	}
}

func (dbm *databaseModels) GetRawRelationalModels() (str string, err error) {
	if err = dbm.populateModelsMap(); err != nil {
		return "", err
	}
	str, err = dbm.marshalMap(dbm.tables)
	if err != nil {
		return "", err
	}
	return str, nil
}

func (dbm *databaseModels) GetRawCollections() (str string, err error) {
	if err = dbm.populateModelsMap(); err != nil {
		return "", err
	}
	str, err = dbm.marshalMap(dbm.collections)
	if err != nil {
		return "", err
	}
	return str, nil
}

func (dbm *databaseModels) GetRawKeyValueModels() (str interface{}, err error) {
	if err = dbm.populateModelsMap(); err != nil {
		return "", err
	}
	str, err = dbm.marshalMap(dbm.keyValues)
	if err != nil {
		return "", err
	}
	return str, nil
}

func (dbm *databaseModels) populateModelsMap() error {
	if len(dbm.modelsMap) != 0 {
		return nil
	}
	var modelsMap map[string]interface{}

	if err := json.Unmarshal([]byte(dbm.rawJSON), &modelsMap); err != nil {
		return err
	}
	dbm.modelsMap = modelsMap
	return nil
}

func (dbm *databaseModels) marshalMap(dbModelName string) (string, error) {
	if value, ok := dbm.modelsMap[dbModelName]; ok {
		bytes, err := json.Marshal(value)
		if err != nil {
			return "", err
		}
		relationalMap := fmt.Sprintf(`{"%v": %v}`, dbModelName, string(bytes))
		return relationalMap, nil
	}
	return "", fmt.Errorf("database model %v does not exist", dbModelName)
}
