package entity

import (
	"github.com/google/uuid"
)

type User struct {
	ID   *uuid.UUID `bun:"type:uuid,default:uuid_generate_v4()"`
	Name string
	Pwd  string
}
