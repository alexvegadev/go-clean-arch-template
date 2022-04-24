package db

import (
	"database/sql"
	"reflect"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestNew(t *testing.T) {
	type args struct {
		user     string
		password string
		host     string
		port     string
		database string
	}
	tests := []struct {
		name    string
		args    args
		want    *sql.DB
		wantErr bool
	}{

		{
			name: "Test Connection to DB",
			args: args{
				user:     "root",
				password: "123",
				host:     "localhost",
				port:     "8080",
				database: "db",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.user, tt.args.password, tt.args.host, tt.args.port, tt.args.database)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
