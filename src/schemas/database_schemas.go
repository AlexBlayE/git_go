package schemas

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Name     string
	Password string
	Token    string
}

type Repository struct {
	Name  string
	Users []struct {
		UserId primitive.ObjectID
		Role   string
	}
}
