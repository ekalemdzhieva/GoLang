package querybuilder

import (
	"fmt"
	"strings"
)

func CreateTable(tableName string, columns map[string]string) string {
	var cols []string
	for col, colType := range columns {
		cols = append(cols, fmt.Sprintf("%s %s", col, colType))
	}
	columnsDef := strings.Join(cols, ", ")
	return fmt.Sprintf("CREATE TABLE %s (%s);", tableName, columnsDef)
}

func SelectQuery(tableName string, columns []string, where map[string]string) string {
	cols := strings.Join(columns, ", ")
	var whereClauses []string
	for col, val := range where {
		whereClauses = append(whereClauses, fmt.Sprintf("%s='%s'", col, val))
	}
	whereClause := strings.Join(whereClauses, " AND ")
	return fmt.Sprintf("SELECT %s FROM %s WHERE %s;", cols, tableName, whereClause)
}

func Insert(tableName string, data map[string]string) string {
	var cols, vals []string
	for col, val := range data {
		cols = append(cols, col)
		vals = append(vals, fmt.Sprintf("'%s'", val))
	}
	columns := strings.Join(cols, ", ")
	values := strings.Join(vals, ", ")
	return fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s);", tableName, columns, values)
}

func Update(tableName string, data map[string]string, where map[string]string) string {
	var setClauses, whereClauses []string
	for col, val := range data {
		setClauses = append(setClauses, fmt.Sprintf("%s='%s'", col, val))
	}
	for col, val := range where {
		whereClauses = append(whereClauses, fmt.Sprintf("%s='%s'", col, val))
	}
	setClause := strings.Join(setClauses, ", ")
	whereClause := strings.Join(whereClauses, " AND ")
	return fmt.Sprintf("UPDATE %s SET %s WHERE %s;", tableName, setClause, whereClause)
}
