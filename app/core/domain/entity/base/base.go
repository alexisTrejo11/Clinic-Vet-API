// Package base contains the base schema to apply in each entity
package base

import (
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
)

type Entity struct {
	id        valueobject.IntegerID
	createdAt time.Time
	updatedAt time.Time
	version   int
}

// NewEntity creates a new base entity
func NewEntity(id valueobject.IntegerID) Entity {
	now := time.Now()
	return Entity{
		id:        id,
		createdAt: now,
		updatedAt: now,
		version:   1,
	}
}

// Getters
func (e Entity) ID() valueobject.IntegerID {
	return e.id
}

func (e Entity) CreatedAt() time.Time {
	return e.createdAt
}

func (e Entity) UpdatedAt() time.Time {
	return e.updatedAt
}

func (e Entity) Version() int {
	return e.version
}

// IncrementVersion updates the version and updatedAt timestamp
func (e *Entity) IncrementVersion() {
	e.version++
	e.updatedAt = time.Now()
}
