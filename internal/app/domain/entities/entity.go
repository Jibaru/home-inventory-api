package entities

type Entity interface {
	EntityID() string
	EntityName() string
}

type IdentifiableEntity struct {
	ID string
}

func NewIdentifiableEntity(id string) *IdentifiableEntity {
	return &IdentifiableEntity{id}
}

func (e *IdentifiableEntity) EntityID() string {
	return e.ID
}

func (e *IdentifiableEntity) EntityName() string {
	return "mock_entity"
}
