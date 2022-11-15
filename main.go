package main

import (
	"fmt"
	"github.com/Hesam-Eskandari/transformer/pkg/repository"
)

func main() {
	jsn := `{"tableName": "orders", "schema": [{"name": "id", "type": "int64"},{"name": "customer_id", "type": "string"}]}`
	relational := repository.NewRelationalModel()
	err := relational.Unmarshal(jsn)
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
