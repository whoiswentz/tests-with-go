package psql

import (
	"database/sql"
	"reflect"
	"testing"
)

func TestUserStore_Find(t *testing.T) {
	type fields struct {
		sql interface {
			Exec(query string, args ...interface{}) (sql.Result, error)
			QueryRow(query string, args ...interface{}) *sql.Row
		}
	}
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			us := &UserStore{
				sql: tt.fields.sql,
			}
			got, err := us.Find(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserStore.Find() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserStore.Find() = %v, want %v", got, tt.want)
			}
		})
	}
}
