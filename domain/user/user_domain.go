package domain

type User struct {
	Name     string `json:"name" bson:"name"`
	LastName string `json:"last_name" bson:"last_name"`
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
	Pin      *uint8 `json:"pin" bson:"pin"`
}
