package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID           primitive.ObjectID `bson:"_id"`
	FirstName    *string            `json:"first_name" validate:"required, min=2, max=100"`
	LastName     *string            `json:"last_name" validate:"required, min=2, max=100"`
	Password     *string            `json:"password" validate:"required, min=8"`
	Email        *string            `json:"email" validate:"required"`
	Contact      *string            `json:"contact" validate:"required, min=10"`
	Token        *string            `json:"token"`
	UserType     *string            `json:"userType" validate:"requied, eq=ADMIN|eq=USER"`
	RefreshToken *string            `json:"refreshToken"`
	Created_at   time.Time          `json:"created_at"`
	Updated_at   time.Time          `json:"updated_at"`
	UserID       *string            `json:"userID"`
}
