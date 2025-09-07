// Package base contains the base schema to apply in each entity
package base

import (
	"time"
)

type Entity[T any] struct {
	id        T
	createdAt time.Time
	updatedAt time.Time
	version   int
}

// NewEntity creates a new base entity
func NewEntity[T any](id T) Entity[T] {
	now := time.Now()
	return Entity[T]{
		id:        id,
		createdAt: now,
		updatedAt: now,
		version:   1,
	}
}

func (e Entity[T]) ID() T {
	return e.id
}

func (e Entity[T]) CreatedAt() time.Time {
	return e.createdAt
}

func (e Entity[T]) UpdatedAt() time.Time {
	return e.updatedAt
}

func (e Entity[T]) Version() int {
	return e.version
}

// IncrementVersion updates the version and updatedAt timestamp
func (e *Entity[T]) IncrementVersion() {
	e.version++
	e.updatedAt = time.Now()
}
