package main

import (
	"testing"

	"github.com/jmoiron/sqlx"
)

func TestSqlxHelper_AddEntity(t *testing.T) {
	type fields struct {
		db             *sqlx.DB
		columns        map[string][]string
		managedColumns []string
	}
	type args struct {
		tagName  string
		entities []interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sh := &SqlxHelper{
				db:             tt.fields.db,
				columns:        tt.fields.columns,
				managedColumns: tt.fields.managedColumns,
			}
			if err := sh.AddEntity(tt.args.tagName, tt.args.entities...); (err != nil) != tt.wantErr {
				t.Errorf("SqlxHelper.AddEntity() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
