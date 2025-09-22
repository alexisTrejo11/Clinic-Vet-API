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
func NewEntity[T any](id T, createdAt, updatedAt time.Time, version int) Entity[T] {
	return Entity[T]{
		id:        id,
		createdAt: createdAt,
		updatedAt: updatedAt,
		version:   version,
	}
}

func CreateEntity[T any](id T) Entity[T] {
	return Entity[T]{
		id:        id,
		createdAt: time.Now(),
		updatedAt: time.Now(),
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

func (e *Entity[T]) SetTimeStamps(createAt, updateAt time.Time) {
	e.createdAt = createAt
	e.updatedAt = updateAt
}

func (e *Entity[T]) SetID(id T) {
	e.id = id
}
