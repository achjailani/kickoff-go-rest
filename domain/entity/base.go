package entity

// Entity is a set of functions that should be implemented by entity
type Entity interface {
	TableName() string
	FilterableFields() []interface{}
	TimeFields() []interface{}
}
