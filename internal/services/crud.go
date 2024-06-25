package services

type CRUD interface {
	CreateTable() error
	Insert(interface{}) error
	Update(interface{}) error
	Delete(interface{}) error
}
