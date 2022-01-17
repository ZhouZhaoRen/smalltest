package model

type StudentMongo struct {
	Id int `bson:"_id"`
	Name string `bson:"name"`
	Gender int `bson:"gender"`
	Country string `bson:"country"`
}
