package main

import (
	"time"

	"github.com/google/uuid"
	"github.com/nikshit-chauhan/rssagg/internal/database"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
}

func databaseUserToUser(dbUser database.User) User {
	return User{
		ID:        dbUser.ID,
		CreatedAt: dbUser.CreateAt,
		UpdatedAt: dbUser.UpdateAt,
		Name:      dbUser.Name,
	}
}
