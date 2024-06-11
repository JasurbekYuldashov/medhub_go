package models

import (
	"time"
)

type Base struct {
	ID        int        `db:"id" json:"id"`
	CreateAt  time.Time  `db:"created_at" json:"createdAt"`
	UpdatedAt time.Time  `db:"updated_at" json:"updatedAt"`
	DeletedAt *time.Time `db:"deleted_at" json:"deletedAt"`
}
