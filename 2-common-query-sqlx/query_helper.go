package main

import (
	"reflect"

	"github.com/jmoiron/sqlx"

	"github.com/knq/snaker"
)

type SqlxHelper struct {
	db             *sqlx.DB
	columns        map[string][]string
	managedColumns []string
}

func (sh *SqlxHelper) AddEntity(tagName string, entities ...interface{}) error {

	for _, entity := range entities {
		cols, name := GetColumns(entity, tagName)
		if _, ok := sh.columns[name]; !ok {
			sh.columns[name] = cols
		}
	}
	return nil
}

func (sh *SqlxHelper) FindOne(id int64, tableName string, out interface{}) error {
	return nil
}

func (sh *SqlxHelper) Find(tableName string, query map[string]interface{}, out interface{}) error {
	return nil
}

func (sh *SqlxHelper) Delete(tableName string, query map[string]interface{}) error {
	return nil
}

func (sh *SqlxHelper) Create(tableName string, data map[string]interface{}) (int64, error) {
	return 0, nil
}

func (sh *SqlxHelper) Update(tableName string, id int64, query, data map[string]interface{}) (int64, error) {
	return 0, nil
}
func New(db *sqlx.DB) *SqlxHelper {

	return &SqlxHelper{
		db:             db,
		columns:        map[string][]string{},
		managedColumns: []string{},
	}
}

// GetColumns from tag name of the struct
func GetColumns(entity interface{}, tagName string) ([]string, string) {
	st := reflect.TypeOf(entity)
	name := st.Name()
	columns := []string{}
	for i := 0; i < st.NumField(); i++ {
		field := st.Field(i)
		if tag, ok := field.Tag.Lookup(tagName); ok {
			if tag == "" {
				columns = append(columns, snaker.CamelToSnake(field.Name))
			} else {
				columns = append(columns, tag)
			}
		}
	}

	return columns, name
}
