package repository

import "encoding/json"

type column struct {
	ColumnName string `json:"name"`
	ColumnType string `json:"type"`
}

type table struct {
	Schema    []column `json:"schema"`
	TableName string   `json:"tableName"`
}

type relationalModel struct {
	Tables []table `json:"tables"`
}

func NewRelationalModel() RelationalModel {
	return &relationalModel{}
}

func (r *relationalModel) Marshal() (string, error) {
	bytes, err := json.Marshal(r)
	return string(bytes), err
}

func (r *relationalModel) Unmarshal(str string) error {
	return json.Unmarshal([]byte(str), &r)
}
