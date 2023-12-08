package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id       primitive.ObjectID `bson:"_id, omitempty"`
	Name     string             `json:"name" bson:"name"`
	LastName string             `json:"last_name" bson:"last_name"`
	Email    string             `json:"email" bson:"email"`
	Password string             `json:"password" bson:"password"`
	Pin      *uint8             `json:"pin" bson:"pin"`
}

type UserSignIn struct {
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}

type Token struct {
	AccessToken  *string `json:"access_token"`
	RefreshToken *string `json:"refresh_token"`
}
