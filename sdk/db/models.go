package db

type User struct {
	ID   string `json:"_id" bson:"_id"`
	Name string `json:"name" bson:"name"`
}
