package main

import (
	"fmt"
	"github.com/Hesam-Eskandari/transformer/pkg/repository"
)

func main() {
	jsn := `{"tables":[{"tableName": "orders", "schema": [{"name": "id", "type": "int64"},{"name": "customer_id", "type": "string"}]}]}`
	//jsnByte, err := os.ReadFile("test.json")
	//jsn = string(jsnByte)
	dbModels := repository.NewDatabaseModels(jsn)
	relationalModel, err := dbModels.GetRawRelationalModels()
	if err != nil {
		fmt.Printf("error: %v", err)
	}
	fmt.Println(relationalModel)
	//configPath := library.GetConfigPath("orders")
	//fmt.Println(configPath.GetInputPath())
	relational := repository.NewRelationalModel()
	err = relational.Unmarshal(relationalModel)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("%+v \n", relational)

	str, err := relational.Marshal()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf(str)
}
