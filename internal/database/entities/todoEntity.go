package entities

import "time"

type TodoEntity struct {
	Id        *string    `bson:"_id"`
	Title     string     `bson:"title"`
	IsActive  bool       `bson:"isActive"`
	CreatedAt time.Time  `bson:"createdAt"`
	UpdatedAt *time.Time `bson:"updatedAt"`
}
