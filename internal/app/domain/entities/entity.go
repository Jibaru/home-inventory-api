package entities

type Entity interface {
	EntityID() string
	EntityName() string
}
