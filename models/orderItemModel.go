package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderItem struct {
	ID 			primitive.ObjectID `bson:"_id"`
	OrderItem_id *string			`json:"orderItem_id"`
	Order_id *string			`json:"order_id"`
	Food_id *string			`json:"food_id"`
	Quantity *int			`json:"quantity"`
	Unit_price *float64			`json:"unit_price"`
	Created_at time.Time		 `json:"created_at"`
	Updated_at time.Time		 `json:"updated_at"`
}