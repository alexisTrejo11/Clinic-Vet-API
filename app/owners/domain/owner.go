package ownerDomain

import "time"

type Owner struct {
	ID        int        `json:"id" db:"id"`
	Photo     *string    `json:"photo,omitempty" db:"photo"` // Opcional
	Name      string     `json:"name" db:"name"`
	LastName  string     `json:"last_name" db:"last_name"`
	UserID    *int       `json:"user_id,omitempty" db:"user_id"`   // Opcional, FK a users.id
	Birthday  *time.Time `json:"birthday,omitempty" db:"birthday"` // Opcional
	Genre     *string    `json:"genre,omitempty" db:"genre"`       // Opcional, "male", "female", "other"
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt time.Time  `json:"updated_at" db:"updated_at"`
}
