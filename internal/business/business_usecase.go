package business

// This is just a example

import "database/sql"

type BusinessUseCaseCRUD interface {
	Create()
	Read()
	Update()
	Delete()
}

type businessUseCaseCRUD struct {
	db *sql.DB
}

func NewBusinessUseCase(db *sql.DB) *businessUseCaseCRUD {
	return &businessUseCaseCRUD{db}
}

func (b businessUseCaseCRUD) Create() {

}

func (b businessUseCaseCRUD) Read() {

}

func (b businessUseCaseCRUD) Update() {

}

func (b businessUseCaseCRUD) Delete() {

}
