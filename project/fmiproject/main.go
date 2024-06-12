package main

import (
	"fmt"
	"sqlbuilder/querybuilder"
)

func main() {

	tableName := "users"
	columns := map[string]string{
		"id":   "INT PRIMARY KEY",
		"name": "VARCHAR(100)",
		"age":  "INT",
	}
	fmt.Println(querybuilder.CreateTable(tableName, columns))

	selectColumns := []string{"id", "name", "age"}
	where := map[string]string{
		"age": "30",
	}
	fmt.Println(querybuilder.SelectQuery(tableName, selectColumns, where))

	insertData := map[string]string{
		"id":   "1",
		"name": "Emanoela Kalemdzhieva",
		"age":  "24",
	}
	fmt.Println(querybuilder.Insert(tableName, insertData))

	updateData := map[string]string{
		"name": "Ralitsa Spasova",
	}
	updateWhere := map[string]string{
		"id": "1",
	}
	fmt.Println(querybuilder.Update(tableName, updateData, updateWhere))
}
