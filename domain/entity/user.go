package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Name     string             `bson:"name"`
	Hash     string             `bson:"hash"`
	Email    string             `bson:"email"`
	Blocked  bool               `bson:"blocked"`
	Username string             `bson:"username"`
	Verified bool               `bson:"verified"`
}
