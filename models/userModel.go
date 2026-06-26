package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID 			primitive.ObjectID `bson:"_id"`
	User_id *string			`json:"user_id"`
	First_name *string			`json:"first_name"`
	Last_name *string			`json:"last_name"`
	Email *string			`json:"email"`	
	Password *string			`json:"password"`
	Avatar	*string			`json:"avatar"`
	Token *string			`json:"token"`
	Refresh_token *string			`json:"refresh_token"`
	Created_at time.Time		 `json:"created_at"`
	Updated_at time.Time		 `json:"updated_at"`
}