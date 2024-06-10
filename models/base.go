package models

type Base struct {
	ID        int `db:"id" json:"id"`
	CreateAt  int `db:"created_at" json:"createdAt"`
	UpdatedAt int `db:"updated_at" json:"updatedAt"`
	DeletedAt int `db:"deleted_at" json:"deletedAt"`
}
