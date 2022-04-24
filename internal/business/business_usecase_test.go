package business

import (
	"database/sql"
	"testing"
)

func Test_businessUseCaseCRUD_Create(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := businessUseCaseCRUD{
				db: tt.fields.db,
			}
			b.Create()
		})
	}
}

func Test_businessUseCaseCRUD_Read(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := businessUseCaseCRUD{
				db: tt.fields.db,
			}
			b.Read()
		})
	}
}

func Test_businessUseCaseCRUD_Update(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := businessUseCaseCRUD{
				db: tt.fields.db,
			}
			b.Update()
		})
	}
}

func Test_businessUseCaseCRUD_Delete(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := businessUseCaseCRUD{
				db: tt.fields.db,
			}
			b.Delete()
		})
	}
}
